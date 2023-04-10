package logic

import (
	"context"
	"fmt"
	"strings"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContainerLogic {
	return &CreateContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContainerLogic) CreateContainer(req *types.CreateContainerRequest) error {
	var worker *types.Node
	var err error

	if req.Container.Metadata.Namespace == "" {
		return fmt.Errorf("container's namespace cannot be empty")
	}

	// 判断 container 的 Namespace 是否存在
	isExist, err := etcdutil.IsExistNamespace(l.svcCtx.Etcd, l.ctx, req.Container.Metadata.Namespace)
	if err != nil {
		return err
	}

	if !isExist {
		return fmt.Errorf("namespace %s does not exist", req.Container.Metadata.Namespace)
	}

	// 判断容器是否已经存在
	key := etcdutil.GenerateKey("container", req.Container.Metadata.Namespace, req.Container.Metadata.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	if found {
		return fmt.Errorf("container %s already exist", req.Container.Metadata.Name)
	}

	// 如果有指定结点，根据选择的结点创建容器
	if req.Container.ContainerConfig.NodeName != "" {
		w, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, req.Container.ContainerConfig.NodeName)
		if err != nil {
			return err
		}

		if !found {
			return fmt.Errorf("node %s does not exist", req.Container.ContainerConfig.NodeName)
		}

		if w.Spec.Unschedulable {
			return fmt.Errorf("the node %s is unschedulable", req.Container.ContainerConfig.NodeName)
		}

		if !w.Status.Working {
			return fmt.Errorf("the node %s is not active", req.Container.ContainerConfig.NodeName)
		}

		if !lo.Contains(w.Roles, "worker") {
			return fmt.Errorf("the node %s is not a worker", req.Container.ContainerConfig.NodeName)
		}
		
		worker = new(types.Node)
		*worker = types.Node{
			Metadata: types.Metadata{
				Namespace: w.Metadata.Namespace,
				Kind:      w.Metadata.Kind,
				Name:      w.Metadata.Name,
			},
			Roles: w.Roles,
			BaseURL: types.NodeURL{
				WorkerURL: w.BaseURL.WorkerURL,
				MasterURL: w.BaseURL.MasterURL,
			},
		}

	} else {
		// 从 etcd 中获取需要创建容器的 worker 结点，根据在线调度算法自动获取
		worker, err = l.svcCtx.Worker(&req.Container)
		if err != nil {
			//return fmt.Errorf("not found worker can run")
			return err
		}
	}

	cli := client.NewClient(worker.BaseURL.WorkerURL)
	c := types.Container{}
	c.Metadata = req.Container.Metadata
	c.Metadata.Kind = "container"
	c.ContainerConfig = req.Container.ContainerConfig
	c.ContainerStatus.Node = worker.Metadata.Name

	if c.Metadata.Name == "" {
		c.Metadata.Name = strings.ReplaceAll(c.ContainerConfig.Image+uuid.New().String(), "-", "")
	}

	expose := make([]client.ExposedPort, 0)
	for _, e := range c.ContainerConfig.Expose {
		expose = append(expose, client.ExposedPort{
			Port:     e.Port,
			Protocol: e.Protocol,
			HostPort: e.HostPort,
		})
	}

	// 访问 worker 结点并创建容器
	res, err := cli.Container.Create(l.ctx, client.CreateContainerRequest{
		ContainerName: c.Metadata.Name,
		Config: client.ContainerConfig{
			Image:    c.ContainerConfig.Image,
			NodeName: c.ContainerConfig.NodeName,
			Command:  c.ContainerConfig.Command,
			Args:     c.ContainerConfig.Args,
			Expose:   expose,
			Env:      c.ContainerConfig.Env,
			Limit:    client.ContainerLimit{
				CPU: c.ContainerConfig.Limit.CPU,
				Memory: c.ContainerConfig.Limit.Memory,
			},
			Request:    client.ContainerRequest{
				CPU: c.ContainerConfig.Request.CPU,
				Memory: c.ContainerConfig.Request.Memory,
			},
		},
	})

	if err != nil {
		return err
	}
	c.ContainerStatus.ContainerID = res.ID
	// 将容器信息写入etcd
	return etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, c)
}

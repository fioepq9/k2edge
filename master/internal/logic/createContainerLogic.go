package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/google/uuid"
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
	// 如果有指定结点，根据选择的结点创建容器
	if req.Container.ContainerConfig.NodeName != "" {
		// Node 的 Namespace
		nodes , err := etcdutil.GetOne[[]types.Node](l.svcCtx.Etcd, l.ctx, "/nodes")
		if err != nil {
			return err
		}

		// 判断结点是否存在
		found := false
		for _, n := range *nodes {
			if n.Metadata.Namespace == req.Container.ContainerConfig.NodeNamespace && n.Metadata.Name == req.Container.ContainerConfig.NodeName && n.Status == "active" {
				worker = new(types.Node)
				*worker =  n
				found = true
				break
			}
		}

		// 未找到结点
		if !found {
			return fmt.Errorf("cannot find node %s", req.Container.ContainerConfig.NodeName)
		}

	} else {
		// 从 etcd 中获取需要创建容器的 worker 结点，根据在线调度算法自动获取
		worker, err = l.svcCtx.Worker()
		if err != nil {
			return fmt.Errorf("not found worker can run")
		}
	}

	cli := client.NewClient(worker.BaseURL.WorkerURL)
	var c types.Container
	c.Metadata = req.Container.Metadata
	c.Metadata.Kind = "container"
	c.ContainerConfig = req.Container.ContainerConfig
	c.ContainerStatus.Node = worker.Metadata.Name

	if c.Metadata.Name == "" {
		c.Metadata.Name = c.ContainerConfig.Image + uuid.New().String()
	}

	fmt.Println("aaaaa")
	// 判断容器是否已经存在
	isExist, err := etcdutil.IsExist(l.svcCtx.Etcd, l.ctx, "/containers", etcdutil.Metadata{
		Namespace: c.Metadata.Namespace,
		Kind: c.Metadata.Kind,
		Name: c.Metadata.Name,
	})

	if err != nil {
		return err
	}

	if isExist {
		return fmt.Errorf("container %s already exist", c.Metadata.Name)
	}

	fmt.Println("bbbbb")
	expose := make([]client.ExposedPort, 0)
	for _, e := range c.ContainerConfig.Expose {
		expose = append(expose, client.ExposedPort{
			Port:     e.Port,
			Protocol: e.Protocol,
			HostPort: e.HostPort,
		})
	}

	// 访问 worker 结点并创建容器
	res, err := cli.Containers().Create(l.ctx, client.CreateContainerRequest{
		ContainerName: req.Container.Metadata.Name,
		Config: client.ContainerConfig{
			Image:   c.ContainerConfig.Image,
			NodeName:    c.ContainerConfig.NodeName,
			Command: c.ContainerConfig.Command,
			Args:    c.ContainerConfig.Args,
			Expose:  expose,
			Env:     c.ContainerConfig.Env,
		},
	})
	fmt.Println(res)
	if err != nil {
		return err
	}
	c.ContainerStatus.ContainerID = res.ID
	fmt.Println(res.ID)
	// 将容器信息写入etcd
	return etcdutil.AddOne(l.svcCtx.Etcd, l.ctx, "/containers", c)
}

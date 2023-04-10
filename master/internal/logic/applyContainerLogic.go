package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyContainerLogic {
	return &ApplyContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyContainerLogic) ApplyContainer(req *types.ApplyContainerRequest) error {
	// 判断容器是否已经存在
	key := etcdutil.GenerateKey("container", req.Container.Metadata.Namespace, req.Container.Metadata.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("container %s does not exist", req.Container.Metadata.Name)
	}
	
	// 获取目的容器在etcd中的信息以便备份恢复
	getFunc := NewGetContainerLogic(l.ctx, l.svcCtx)
	config, err := getFunc.GetContainer(&types.GetContainerRequest{
		Namespace: req.Container.Metadata.Namespace,
		Name: req.Container.Metadata.Name,
	})

	// 备份失败，直接终止
	if err != nil {
		return fmt.Errorf("backup container %s's configuration failed", req.Container.Metadata.Name)
	}

	// 创建容器 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.
	var worker *types.Node

	if req.Container.Metadata.Namespace == "" {
		return fmt.Errorf("container's namespace cannot be empty")
	}
	if req.Container.Metadata.Name == "" {
		return fmt.Errorf("container's name cannot be empty")
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
			return fmt.Errorf("the node %s is unschedulable", req.Container.Metadata.Name)
		}

		if !w.Status.Working {
			return fmt.Errorf("the node %s is not active", req.Container.Metadata.Name)
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
			return err
		}
	}

	cli := client.NewClient(worker.BaseURL.WorkerURL)
	var c types.Container
	c.Metadata = req.Container.Metadata
	c.Metadata.Kind = "container"
	c.ContainerConfig = req.Container.ContainerConfig
	c.ContainerStatus.Node = worker.Metadata.Name

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
	// 在worker节点创建容器成功 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// 将容器信息写入etcd
	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, c)
	if err != nil { // 写入etcd失败
		// 删除worker中新创建的容器
		errd := cli.Container.Stop(l.ctx, client.StopContainerRequest{
			ID:      c.ContainerStatus.ContainerID,
		})
	
		if errd != nil {
			return fmt.Errorf("apply container failed, recover failed : %s", err.Error())
		}
	
		errd = cli.Container.Remove(l.ctx, client.RemoveContainerRequest{
			ID:            c.ContainerStatus.ContainerID,
		})
	
		if errd != nil {
			return fmt.Errorf("apply container failed, recover failed : %s", err.Error())
		}
		
		return err
	}

	// 写入etcd成功后，删除原本的container
	// 删除worker中新创建的容器
	err = cli.Container.Stop(l.ctx, client.StopContainerRequest{
		ID:      config.Container.ContainerStatus.ContainerID,
	})

	if err != nil {
		return fmt.Errorf("apply container failed, delete orignal container failed : %s", err.Error())
	}

	err = cli.Container.Remove(l.ctx, client.RemoveContainerRequest{
		ID:            config.Container.ContainerStatus.ContainerID,
	})

	if err != nil {
		return fmt.Errorf("apply container failed, delete orignal container failed : %s", err.Error())
	}
	return nil
}
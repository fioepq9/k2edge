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
	// 从 etcd 中获取需要创建容器的 worker 结点
	worker, err := l.svcCtx.Worker()
	if err != nil {
		return fmt.Errorf("not found worker can run")
	}
	cli := client.NewClient(worker.BaseURL)
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
	res, err := cli.Containers().Create(l.ctx, client.CreateContainerRequest{
		ContainerName: req.Container.Metadata.Name,
		Config: client.ContainerConfig{
			Image:   c.ContainerConfig.Image,
			Node:    c.ContainerConfig.Node,
			Command: c.ContainerConfig.Command,
			Args:    c.ContainerConfig.Args,
			Expose:  expose,
			Env:     c.ContainerConfig.Env,
		},
	})
	if err != nil {
		return err
	}
	c.ContainerStatus.ContainerID = res.ID
	return etcdutil.AddOne(l.svcCtx.Etcd, l.ctx, "/containers", c)
}

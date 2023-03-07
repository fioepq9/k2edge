package logic

import (
	"context"
	"fmt"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/master/model"
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
	worker, err := l.svcCtx.Worker()
	if err != nil {
		return fmt.Errorf("not found worker can run")
	}
	cli := client.NewClient(worker.BaseURL)
	var container model.Container
	container.Container = req.Container
	container.ContainerStatus = model.ContainerStatus{
		Node:          worker.Metadata.Name,
		ContainerName: fmt.Sprintf("%s-%d", req.Container.Metadata.Name, time.Now().Unix()),
	}
	r, err := cli.Containers().Run(l.ctx, client.RunContainerRequest{
		ContainerName: container.ContainerStatus.ContainerName,
		Config: client.ContainerConfig{
			Image: container.ContainerTemplate.Image,
			Cmd:   append([]string{container.ContainerTemplate.Command}, container.ContainerTemplate.Args...),
			Env:   container.ContainerTemplate.Env,
		},
	})
	if err != nil {
		return err
	}
	container.ContainerStatus.ContainerID = r.ID
	return etcdutil.AddOne(l.svcCtx.Etcd, l.ctx, "/containers", container)
}

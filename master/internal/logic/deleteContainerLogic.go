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

type DeleteContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContainerLogic {
	return &DeleteContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContainerLogic) DeleteContainer(req *types.DeleteContainerRequest) error {
	key := "/containers"
	//根据 container 里 nodeName 去 etcd 里查询的 nodeBaseURL
	containers, err := etcdutil.GetOne[[]types.Container](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	// 判断 container 是否存在, 存在则获取 container 信息
	var container types.Container
	found := false
	for _, container = range *containers {
		if container.Metadata.Namespace == req.Namespace && container.Metadata.Name == req.Name {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("container %s does not exist", req.Name)
	}

	// 获取 node 的 BaseURL
	worker, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, container.ContainerStatus.Node)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("cannot find container %s info", req.Name)
	}
	// 向特定的 worker 结点发送获取conatiner信息的请求
	cli := client.NewClient(worker.BaseURL.WorkerURL)
	err = cli.Containers().Stop(l.ctx, client.StopContainerRequest{
		ID: container.ContainerStatus.ContainerID,
		Timeout: req.Timeout,
	})

	if err != nil {
		return  err
	}


	err = cli.Containers().Remove(l.ctx, client.RemoveContainerRequest{
		ID:            container.ContainerStatus.ContainerID,
		RemoveVolumes: req.RemoveVolumnes,
		RemoveLinks:   req.RemoveLinks,
		Force:         req.Force,
	})

	if err != nil {
		return err
	}

	err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key, func(item types.Container, index int) bool {
		return item.Metadata.Namespace == req.Namespace && item.Metadata.Name == req.Name && item.Metadata.Kind == "container"
	})

	if err != nil {
		return err
	}

	return nil
}

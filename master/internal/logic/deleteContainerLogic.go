package logic

import (
	"context"
	"fmt"
	"time"

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
	key := etcdutil.GenerateKey("container", req.Namespace, req.Name)

	// 判断 container 是否存在, 存在则获取 container 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("container %s does not exist", req.Name)
	}

	container, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}
	c := (*container)[0]

	node, err := etcdutil.GetOne[types.Node](l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("node", etcdutil.SystemNamespace, c.ContainerStatus.Node))

	if err != nil {
		return fmt.Errorf("cannot get node %s info", c.ContainerStatus.Node)
	}

	n := (*node)[0]
	// 获取 node 的 BaseURL
	worker, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, n.Metadata.Name)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("cannot find container %s info", req.Name)
	}

	if !worker.Status.Working {
		return fmt.Errorf("the node where the container is located is not active")
	}

	// 向特定的 worker 结点发送获取conatiner信息的请求
	cli := client.NewClient(worker.BaseURL.WorkerURL)
	err = cli.Container.Stop(l.ctx, client.StopContainerRequest{
		ID:      c.ContainerStatus.ContainerID,
		Timeout: req.Timeout * int(time.Second),
	})

	if err != nil {
		return err
	}

	err = cli.Container.Remove(l.ctx, client.RemoveContainerRequest{
		ID:            c.ContainerStatus.ContainerID,
		RemoveVolumes: req.RemoveVolumnes,
		RemoveLinks:   req.RemoveLinks,
		Force:         req.Force,
	})

	if err != nil {
		return err
	}

	err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key)

	if err != nil {
		return err
	}
	err = etcdutil.NodeDeleteRequest(l.svcCtx.Etcd, l.ctx, worker.Metadata.Name, c.ContainerConfig.Request.CPU, c.ContainerConfig.Request.Memory)
	if err != nil {
		return err
	}

	return nil
}

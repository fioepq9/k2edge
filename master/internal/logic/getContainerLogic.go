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

type GetContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContainerLogic {
	return &GetContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContainerLogic) GetContainer(req *types.GetContainerRequest) (resp *types.GetContainerResponse, err error) {
	key := etcdutil.GenerateKey("container", req.Namespace, req.Name)
	// 判断 container 是否存在, 存在则获取 container 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("container %s does not exist", req.Name)
	}

	//根据 container 里 nodeName 去 etcd 里查询的 nodeBaseURL
	containers, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	container := (*containers)[0]
	// 获取 node 的 BaseURL
	worker, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, container.ContainerStatus.Node)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("cannot find container %s info", req.Name)
	}

	if !worker.Status.Working {
		return nil, fmt.Errorf("the node where the container is located is not active")
	}

	// 向特定的 work 结点发送获取conatiner信息的请求
	cli := client.NewClient(worker.BaseURL.WorkerURL)
	containerInfo, err := cli.Container.Status(l.ctx, client.ContainerStatusRequest{
		ID: container.ContainerStatus.ContainerID,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.GetContainerResponse)
	resp.Container = container
	resp.Container.ContainerStatus.Info = containerInfo.Status

	return resp, nil
}

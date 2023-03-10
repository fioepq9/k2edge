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
	//根据 container 名字去 etcd 里查询所在的 node 
	containers, err := etcdutil.GetOne[[]types.Container](l.svcCtx.Etcd, l.ctx, "/containers")
	if err != nil {
		return nil, err
	}

	var container types.Container
	var nodeName, nodeBaseURL string
	for _, c := range *containers {
		if c.Metadata.Namespace == req.Metadata.Namespace && c.Metadata.Name == req.Metadata.Name{
			nodeName = c.ContainerStatus.Node
			container = c
			break
		}
	}
	
	// 根据 node 名字查询 worker node 的 BaseURL
	nodes, err := etcdutil.GetOne[[]types.Node](l.svcCtx.Etcd, l.ctx, "/nodes")

	if err != nil {
		return nil, err
	}

	for _, n := range *nodes {
		if n.Metadata.Name == nodeName {
			nodeBaseURL = n.BaseURL.WorkerURL
			break
		}
	}

	if nodeBaseURL == "" || container.ContainerStatus.ContainerID  == ""  {
		return nil, fmt.Errorf("container %s does not exist", req.Metadata.Name)
	}
	
	// 向特定的 work 结点发送获取conatiner信息的请求
	cli := client.NewClient(nodeBaseURL)
	containerInfo, err := cli.Containers().Status(l.ctx, client.ContainerStatusRequest{
		ID: container.ContainerStatus.ContainerID,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.GetContainerResponse)
	resp.Container = container
	resp.Container.ContainerStatus.Info = containerInfo

	return resp, nil
}

package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContainerLogic {
	return &ListContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContainerLogic) ListContainer(req *types.ListContainerRequest) (resp *types.ListContainerResponse, err error) {
	resp = new(types.ListContainerResponse)

	found, err := etcdutil.IsExistNamespace(l.svcCtx.Etcd, l.ctx, req.Namespace)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("namespace %s does not exist", req.Namespace)
	}

	//根据 container 里 nodeName 去 etcd 里查询的 nodeBaseURL 
	containers, err := etcdutil.GetOne[[]types.Container](l.svcCtx.Etcd, l.ctx, "/containers")
	if err != nil {
		return nil, err
	}

	for _, container := range *containers {
		if container.Metadata.Namespace == req.Namespace {
			resp.ContainerSimpleInfo = append(resp.ContainerSimpleInfo, types.ContainerSimpleInfo{
				Name: container.Metadata.Name,
				Namespace: container.Metadata.Namespace,
				Status: container.ContainerStatus.Status,
				Node: container.ContainerStatus.Node,
			})
		}
	}

	return resp, nil
}

package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListContainersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContainersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContainersLogic {
	return &ListContainersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContainersLogic) ListContainers(req *types.ListContainersRequest) (resp *types.ListContainersResponse, err error) {
	containers, err := l.svcCtx.Docker.ContainerList(l.ctx, dtypes.ContainerListOptions{
		Size:   req.Size,
		All:    req.All,
		Latest: req.Latest,
		Since:  req.Since,
		Before: req.Before,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.ListContainersResponse)
	resp.Containers = containers
	return resp, nil
}

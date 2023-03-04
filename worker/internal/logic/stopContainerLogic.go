package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/docker/docker/api/types/container"
	"github.com/zeromicro/go-zero/core/logx"
)

type StopContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopContainerLogic {
	return &StopContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopContainerLogic) StopContainer(req *types.StopContainerRequest) error {
	var timeout *int
	if req.Timeout != 0 {
		timeout = &req.Timeout
	}
	return l.svcCtx.DockerClient.ContainerStop(l.ctx, req.ID, container.StopOptions{
		Timeout: timeout,
	})
}

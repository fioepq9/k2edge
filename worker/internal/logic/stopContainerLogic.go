package logic

import (
	"context"
	"time"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

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
	var timeout *time.Duration
	if req.Timeout != 0 {
		timeout = new(time.Duration)
		*timeout = time.Duration(req.Timeout)
	}
	return l.svcCtx.Docker.ContainerStop(l.ctx, req.ID, timeout)
}

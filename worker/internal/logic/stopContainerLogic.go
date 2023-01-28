package logic

import (
	"context"

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

func (l *StopContainerLogic) StopContainer(req *types.StopContainerRequest) (resp *types.StopContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

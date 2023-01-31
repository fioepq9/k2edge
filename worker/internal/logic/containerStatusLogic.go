package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainerStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContainerStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainerStatusLogic {
	return &ContainerStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContainerStatusLogic) ContainerStatus(req *types.ContainerStatusRequest) (resp *types.ContainerStatusResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

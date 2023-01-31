package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartContainerLogic {
	return &StartContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartContainerLogic) StartContainer(req *types.StartContainerRequest) (resp *types.StartContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunContainerLogic {
	return &RunContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunContainerLogic) RunContainer(req *types.RunContainerRequest) (resp *types.RunContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

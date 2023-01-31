package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainerTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContainerTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainerTopLogic {
	return &ContainerTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContainerTopLogic) ContainerTop(req *types.ContainerTopRequest) (resp *types.ContainerTopResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

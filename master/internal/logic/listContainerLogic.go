package logic

import (
	"context"

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

func (l *ListContainerLogic) ListContainer() (resp *types.ListContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

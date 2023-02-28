package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNamespaceLogic {
	return &ListNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNamespaceLogic) ListNamespace() (resp *types.ListNamespaceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

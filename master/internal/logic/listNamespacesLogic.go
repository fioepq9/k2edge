package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNamespacesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNamespacesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNamespacesLogic {
	return &ListNamespacesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNamespacesLogic) ListNamespaces(req *types.ListNamespacesRequest) (resp *types.ListNamespacesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

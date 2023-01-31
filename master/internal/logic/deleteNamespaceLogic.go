package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNamespaceLogic {
	return &DeleteNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNamespaceLogic) DeleteNamespace(req *types.DeleteNamespaceRequest) (resp *types.DeleteNamespaceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

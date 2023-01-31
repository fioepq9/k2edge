package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamespaceLogic {
	return &GetNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNamespaceLogic) GetNamespace(req *types.GetNamespaceRequest) (resp *types.GetNamespaceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

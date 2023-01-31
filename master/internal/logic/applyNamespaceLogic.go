package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyNamespaceLogic {
	return &ApplyNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyNamespaceLogic) ApplyNamespace(req *types.DeleteNamespaceRequest) (resp *types.DeleteNamespaceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

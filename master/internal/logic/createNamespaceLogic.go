package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNamespaceLogic {
	return &CreateNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNamespaceLogic) CreateNamespace(req *types.CreateNamespaceRequest) (resp *types.CreateNamespaceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

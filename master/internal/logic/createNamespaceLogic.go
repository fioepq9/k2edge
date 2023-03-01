package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/model"

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

func (l *CreateNamespaceLogic) CreateNamespace(req *types.CreateNamespaceRequest) error {
	n := l.svcCtx.DatabaseQuery.Namespace
	namespace := model.Namespace{Name: req.Name, Status: "Active"}

	dbErr := n.WithContext(l.ctx).Create(&namespace)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

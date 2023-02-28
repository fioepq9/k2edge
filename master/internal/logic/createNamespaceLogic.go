package logic

import (
	"context"
	"fmt"
	"time"

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
	fmt.Println("aaaa")
	namespace := model.Namespace{Name: req.Name, Labels: req.Labels, Annotations: req.Annotations, Status: "Active", CreateTime: time.Now()}
	err := l.svcCtx.DatabaseQuery.Namespace.WithContext(l.ctx).Create(&namespace)

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	
	return nil
}

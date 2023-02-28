package logic

import (
	"context"
	"fmt"
	"time"

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

func (l *CreateNamespaceLogic) CreateNamespace(req *types.CreateNamespaceRequest) error {
	namespace := types.Namespace{Name: req.Name, Labels: req.Labels, Annotations: req.Annotations, Status: "Active", CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	result := l.svcCtx.DatabaseClient.Create(&namespace)

	if result.Error != nil {
		return fmt.Errorf(result.Error.Error(), namespace)
	}

	return nil
}

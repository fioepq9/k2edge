package logic

import (
	"context"
	"fmt"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/model"

	"github.com/go-sql-driver/mysql"
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
	err := n.WithContext(l.ctx).Create(&namespace)
	if err != nil {
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				return fmt.Errorf("namespace %s is exist", req.Name)
			}
		}
		return err
	}

	return nil
}

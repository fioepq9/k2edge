package logic

import (
	"context"
	"fmt"

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

func (l *DeleteNamespaceLogic) DeleteNamespace(req *types.DeleteNamespaceRequest) error {
	n := l.svcCtx.DatabaseQuery.Namespace;
	result, dbErr := n.WithContext(l.ctx).Where(n.Name.Eq(req.Name)).Delete()

	if dbErr != nil {
		return dbErr
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("namespace is not exists")
	}
	return nil
}

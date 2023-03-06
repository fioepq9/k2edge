package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNamespaceLogic {
	return &ListNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNamespaceLogic) ListNamespace(req *types.ListNamespaceRequest) (resp *types.ListNamespaceResponse, err error) {
	// n := l.svcCtx.DatabaseQuery.Namespace

	// var namespaces []*model.Namespace
	// var dbErr error
	// if !req.All {
	// 	namespaces, dbErr = n.WithContext(l.ctx).Where(n.Status.Eq("Active")).Find()
	// } else {
	// 	namespaces, dbErr = n.WithContext(l.ctx).Where().Find()
	// }

	// if dbErr != nil {
	// 	return nil, dbErr
	// }

	// resp = new(types.ListNamespaceResponse)
	// for _, namespace := range namespaces {
	// 	resp.Namespaces = append(resp.Namespaces, types.Namespace{
	// 		Name: namespace.Name,
	// 		Status: namespace.Status,
	// 		Age: time.Since(namespace.CreatedTime).Round(time.Second).String(),
	// 	})
	// }

	return resp, nil
}

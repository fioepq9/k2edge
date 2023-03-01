package logic

import (
	"context"
	"time"

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
	n := l.svcCtx.DatabaseQuery.Namespace
	namespace, dbErr := n.WithContext(l.ctx).Where(n.Name.Eq(req.Name)).First()

	if dbErr != nil {
		return nil, dbErr
	}
	resp = new(types.GetNamespaceResponse)
	resp.NamespaceInfo = types.Namespace{
		Name:   namespace.Name,
		Status: namespace.Status,
		Age:    time.Since(namespace.CreatedTime).Round(time.Second).String(),
	}
	return resp, nil
}

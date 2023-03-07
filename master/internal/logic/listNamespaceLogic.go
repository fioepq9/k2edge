package logic

import (
	"context"
	"time"

	"k2edge/etcdutil"
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
	namespace, err := etcdutil.GetOne[[]types.Namespace](l.svcCtx.Etcd, l.ctx, "/namespaces")
	if err != nil {
		return nil, err
	}

	resp = new(types.ListNamespaceResponse)
	for _, n := range *namespace {
		if req.All || n.Status == "Active" { // 判断是否需要返回所有数组
			resp.Namespaces = append(resp.Namespaces, types.GetNamespaceResponse{
				Kind:   n.Kind,
				Name:   n.Name,
				Status: n.Status,
				Age:    time.Since(time.Unix(n.CreateTime, 0)).Round(time.Second).String(),
			})
		}
	}

	return resp, nil
}

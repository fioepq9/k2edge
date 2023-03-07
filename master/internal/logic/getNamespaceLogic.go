package logic

import (
	"context"
	"fmt"
	"time"

	"k2edge/etcdutil"
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
	namespace, err := etcdutil.GetOne[[]types.Namespace](l.svcCtx.Etcd, l.ctx, "/namespace")
	if err != nil {
		return nil, err
	}

	//遇到 name 相同的 namespace 则返回
	resp = new(types.GetNamespaceResponse)
	for _, n := range *namespace {
		if n.Name == req.Name {
			resp.Kind = n.Kind
			resp.Name = n.Name
			resp.Status = n.Status
			resp.Age = time.Since(time.Unix(n.CreateTime, 0)).Round(time.Second).String()
			return resp, nil
		}
	}

	return nil, fmt.Errorf("namespace %s does not exist", req.Name)
}

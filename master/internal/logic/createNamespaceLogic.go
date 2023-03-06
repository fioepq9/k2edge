package logic

import (
	"context"
	"fmt"

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
	namespaces, err := l.svcCtx.Etcd.Get(l.ctx, "aaa")
	if err != nil {
		return err
	}

	for _, KV:= range namespaces.Kvs[0].Value  {
		fmt.Println(string(KV))
	}

	if err != nil {
		return err
	}
	return nil
}

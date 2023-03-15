package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
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
	key := "/namespaces"
	// 判断原本是否存在 namespace
	value, err := etcdutil.GetArray[types.Namespace](l.svcCtx.Etcd, l.ctx, "/namespaces")
	if err != nil {
		return err
	}
	found := false
	for _, n := range *value {
		if n.Name == req.Name{
			found = true
		}
	}

	if !found {
		return fmt.Errorf("namespace %s does exists", req.Name)
	}

	err = etcdutil.DeleteOneValue(l.svcCtx.Etcd, l.ctx, key, func(item types.Namespace, index int) bool {
		return item.Name == req.Name 
	})
	
	if err != nil {
		return  err
	}
	return nil
}

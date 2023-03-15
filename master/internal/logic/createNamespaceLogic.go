package logic

import (
	"context"
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"time"

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
	key := "/namespaces"

	if req.Name == "" {
		return fmt.Errorf("namespace's name cannot be empty")
	}
	// 判断是否已存在 namespace
	value, err := etcdutil.GetArray[types.Namespace](l.svcCtx.Etcd, l.ctx, "/namespaces")
	if err != nil {
		return err
	}
	found := false
	for _, n := range *value {
		if n.Name == req.Name {
			found = true
		}
	}

	if found {
		return fmt.Errorf("namespace %s already exists", req.Name)
	}


	// 插入 namespace
	newNamespace := types.Namespace{
		Kind:       "namespace",
		Name:       req.Name,
		Status:     "active",
		CreateTime: time.Now().Unix(),
	}

	err = etcdutil.AddOneValue(l.svcCtx.Etcd, l.ctx, key, newNamespace)
	if err != nil {
		return err
	}

	return nil
}

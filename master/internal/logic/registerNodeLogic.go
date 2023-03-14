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

type RegisterNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterNodeLogic {
	return &RegisterNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterNodeLogic) RegisterNode(req *types.RegisterRequest) error {
	key := etcdutil.GenerateKey("node", etcdutil.SystemNamespace, req.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)

	if err != nil {
		return err
	}

	if found {
		return fmt.Errorf("node %s already exists", req.Name)
	}

	// 插入 node
	newNode := types.Node{
		Metadata: types.Metadata{
			Namespace: etcdutil.SystemNamespace,
			Kind:      "node",
			Name:      req.Name,
		},
		Roles:        req.Roles,
		BaseURL:      req.BaseURL,
		Status:       "Active",
		RegisterTime: time.Now().Unix(),
	}

	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, newNode)
	if err != nil {
		return err
	}
	return nil
}

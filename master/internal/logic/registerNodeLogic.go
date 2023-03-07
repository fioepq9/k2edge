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
	key := "/nodes"
	isExist, err := etcdutil.IsExist(l.svcCtx.Etcd, l.ctx, key, etcdutil.Metadata{
		Namespace: req.Namespace,
		Kind: "node",
		Name: req.Name,
	})

	if err != nil {
		return err
	}

	if isExist {
		return fmt.Errorf("node %s already exists", req.Name)
	}

	// 插入 node
	newNode := types.Node{
		Metadata : types.Metadata{
			Namespace: req.Namespace,
			Kind: "node",
			Name: req.Name,
		}, 
		Roles: req.Roles,
		BaseURL: req.BaseURL,      
		Status: "Active",     
		RegisterTime: time.Now().Unix(), 
	}

	err = etcdutil.AddOne(l.svcCtx.Etcd, l.ctx, key, newNode)
	if err != nil {
		return err
	}
	return nil
}

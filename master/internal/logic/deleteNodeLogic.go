package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNodeLogic {
	return &DeleteNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNodeLogic) DeleteNode(req *types.DeleteRequest) error {
	key := "/nodes"
	isExist, err := etcdutil.IsExist(l.svcCtx.Etcd, l.ctx, key, etcdutil.Metadata{
		Namespace: req.Metadata.Namespace,
		Kind: req.Metadata.Kind,
		Name: req.Metadata.Name,
	})

	if err != nil {
		return err
	}

	if !isExist {
		return fmt.Errorf("node %s does not exists", req.Metadata.Name)
	}

	err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key, func(item types.Node, index int) bool {
		if item.Metadata.Name == req.Metadata.Name &&
		 item.Metadata.Kind == req.Metadata.Kind && item.Metadata.Namespace == req.Metadata.Namespace{
			return true
		} else {
			return false
		}
	})

	// node, err := etcdutil.GetOne[[]types.Node](l.svcCtx.Etcd, l.ctx, key)
	// if err != nil {
	// 	return err
	// }

	// // 判断是否已存在 node
	// for _, n := range *node {
	// 	if n.Metadata.Name == req.Name {
	// 		return fmt.Errorf("node %s already exists", req.Name)
	// 	}
	// }

	// return nil
	if err != nil {
		return err
	}
	return nil
}

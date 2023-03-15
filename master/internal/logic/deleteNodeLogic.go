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
	key := etcdutil.GenerateKey("node", etcdutil.SystemNamespace, req.Name)

	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)

	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("node %s does not exists", req.Name)
	}

	err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key)

	if err != nil {
		return err
	}
	return nil
}

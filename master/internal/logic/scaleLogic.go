package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScaleLogic {
	return &ScaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScaleLogic) Scale(req *types.ScaleRequest) error {
	// 判断 deployment 是否已经存在
	key := etcdutil.GenerateKey("deployment", req.Namespace, req.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("deployment %s does not exist", req.Name)
	}

	if req.Replicas <= 0 {
		return fmt.Errorf("the replicas of deployment must be more than 0")
	}

	deployments, err := etcdutil.GetOne[types.Deployment](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}
	deployment := (*deployments)[0]
	deployment.Config.Replicas = req.Replicas
	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, deployment)
	if err != nil {
		return err
	}
	return err
}

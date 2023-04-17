package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDrainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrainLogic {
	return &DrainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrainLogic) Drain(req *types.DrainRequest) error {
	key := etcdutil.GenerateKey("node", etcdutil.SystemNamespace, req.Name)

	// 查看 node 存不存在
	node, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, req.Name)

	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("node %s does not exists", req.Name)
	}

	if !node.Status.Working {
		return fmt.Errorf("the node %s is not active", req.Name)
	}

	node.Spec.Unschedulable = true

	/*驱逐所有容器*/
	
	/**/

	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, types.Node{
		Metadata: types.Metadata(node.Metadata),
		Roles: node.Roles,
		BaseURL: types.NodeURL(node.BaseURL),
		Spec: types.Spec{
			Unschedulable: node.Spec.Unschedulable,
			Capacity: types.Capacity(node.Spec.Capacity),
		},
		RegisterTime: node.RegisterTime,
		Status: types.Status{
			Working: node.Status.Working,
			Allocatable: types.Allocatable(node.Status.Allocatable),
			Condition: types.Condition{
				Ready: types.ConditionInfo(node.Status.Condition.Ready),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentLogic {
	return &GetDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeploymentLogic) GetDeployment(req *types.GetDeploymentRequest) (resp *types.GetDeploymentResponse, err error) {
	key := etcdutil.GenerateKey("deployment", req.Namespace, req.Name)
	// 判断 deployment 是否存在, 存在则获取 deployment 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("deployment %s does not exist", req.Name)
	}

	//根据 deployment 里 nodeName 去 etcd 里查询的 nodeBaseURL
	deployment, err := etcdutil.GetOne[types.Deployment](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	resp = new(types.GetDeploymentResponse)
	resp.Deployment = (*deployment)[0]

	return resp, nil
}

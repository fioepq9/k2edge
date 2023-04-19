package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClusterInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterInfoLogic {
	return &ClusterInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterInfoLogic) ClusterInfo() (resp *types.ClusterInfoResponse, err error) {
	resp = new(types.ClusterInfoResponse)
	resp.ClusterInfo = "K2edge v0.0.1\nauthor: fioepq9、tino"
	return resp, nil
}

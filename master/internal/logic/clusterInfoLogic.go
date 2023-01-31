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
	// todo: add your logic here and delete this line

	return
}

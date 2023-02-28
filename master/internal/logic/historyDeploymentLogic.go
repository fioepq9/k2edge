package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryDeploymentLogic {
	return &HistoryDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryDeploymentLogic) HistoryDeployment(req *types.HistoryDeploymentRequest) (resp *types.HistoryDeploymentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

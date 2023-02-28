package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogsDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogsDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogsDeploymentLogic {
	return &LogsDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogsDeploymentLogic) LogsDeployment(req *types.LogsDeploymentRequest) (resp *types.LogsDeploymentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

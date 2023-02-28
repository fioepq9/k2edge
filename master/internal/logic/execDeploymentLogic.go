package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecDeploymentLogic {
	return &ExecDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecDeploymentLogic) ExecDeployment(req *types.ExecDeploymentRequest) (resp *types.ExecDeploymentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

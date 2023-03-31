package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyDeploymentLogic {
	return &ApplyDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyDeploymentLogic) ApplyDeployment(req *types.ApplyDeploymentRequest) error {
	// todo: add your logic here and delete this line

	return nil
}

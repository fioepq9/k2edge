package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}

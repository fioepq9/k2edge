package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttachDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttachDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachDeploymentLogic {
	return &AttachDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttachDeploymentLogic) AttachDeployment(req *types.AttachDeploymentRequest) (resp *types.AttachDeploymentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

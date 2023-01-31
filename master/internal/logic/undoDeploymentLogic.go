package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UndoDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUndoDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UndoDeploymentLogic {
	return &UndoDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UndoDeploymentLogic) UndoDeployment(req *types.UndoDeploymentRequest) (resp *types.UndoDeploymentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

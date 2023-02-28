package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCronJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCronJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCronJobLogic {
	return &DeleteCronJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCronJobLogic) DeleteCronJob(req *types.DeleteCronJobRequest) (resp *types.DeleteCronJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

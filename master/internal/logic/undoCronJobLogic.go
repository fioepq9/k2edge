package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UndoCronJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUndoCronJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UndoCronJobLogic {
	return &UndoCronJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UndoCronJobLogic) UndoCronJob(req *types.UndoCronJobRequest) (resp *types.UndoCronJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

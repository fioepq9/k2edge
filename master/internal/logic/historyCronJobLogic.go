package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryCronJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryCronJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryCronJobLogic {
	return &HistoryCronJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryCronJobLogic) HistoryCronJob(req *types.HistoryCronJobRequest) (resp *types.HistoryCronJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyCronJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyCronJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyCronJobLogic {
	return &ApplyCronJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyCronJobLogic) ApplyCronJob(req *types.ApplyCronJobRequest) (resp *types.ApplyCronJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

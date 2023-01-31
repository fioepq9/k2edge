package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogsCronJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogsCronJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogsCronJobLogic {
	return &LogsCronJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogsCronJobLogic) LogsCronJob(req *types.LogsCronJobRequest) (resp *types.LogsCronJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

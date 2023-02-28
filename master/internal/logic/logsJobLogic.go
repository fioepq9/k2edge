package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogsJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogsJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogsJobLogic {
	return &LogsJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogsJobLogic) LogsJob(req *types.LogsJobRequest) (resp *types.LogsJobResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogsContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogsContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogsContainerLogic {
	return &LogsContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogsContainerLogic) LogsContainer(req *types.LogsContainerRequest) (resp *types.LogsContainerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

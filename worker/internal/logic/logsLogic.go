package logic

import (
	"context"
	"io"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogsLogic {
	return &LogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogsLogic) Logs(req *types.LogsRequest) (io.ReadCloser, error) {
	d := l.svcCtx.Docker
	rd, err := d.ContainerLogs(l.ctx, req.Container, dtypes.ContainerLogsOptions{
		ShowStdout: req.ShowStdout,
		ShowStderr: req.ShowStderr,
		Since:      req.Since,
		Until:      req.Until,
		Timestamps: req.Timestamps,
		Follow:     req.Follow,
		Tail:       req.Tail,
		Details:    req.Details,
	})
	if err != nil {
		return nil, err
	}
	return rd, nil
}

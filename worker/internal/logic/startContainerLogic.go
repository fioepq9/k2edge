package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type StartContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartContainerLogic {
	return &StartContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartContainerLogic) StartContainer(req *types.StartContainerRequest) error {
	return l.svcCtx.Docker.ContainerStart(l.ctx, req.ID, dtypes.ContainerStartOptions{
		CheckpointID:  req.CheckpointID,
		CheckpointDir: req.CheckpointDir,
	})
}

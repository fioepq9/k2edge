package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	typesInternal "k2edge/worker/internal/types"

	"github.com/docker/docker/api/types"
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

func (l *StartContainerLogic) StartContainer(req *typesInternal.StartContainerRequest) error {
	return l.svcCtx.DockerClient.ContainerStart(l.ctx, req.ID, types.ContainerStartOptions{
		CheckpointID:  req.CheckpointID,
		CheckpointDir: req.CheckpointDir,
	})
}

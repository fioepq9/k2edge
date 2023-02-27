package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunContainerLogic {
	return &RunContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunContainerLogic) RunContainer(req *types.RunContainerRequest) error {
	conf := req.Config.DockerFormat()
	_, err := l.svcCtx.DockerClient.ContainerCreate(l.ctx, &conf, nil, nil, nil, req.ContainerName)
	if err != nil {
		return err
	}
	return nil
}

package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	typesInternal "k2edge/worker/internal/types"

	"github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveContainerLogic {
	return &RemoveContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveContainerLogic) RemoveContainer(req *typesInternal.RemoveContainerRequest) error {
	return l.svcCtx.DockerClient.ContainerRemove(l.ctx, req.ID, types.ContainerRemoveOptions{
		RemoveVolumes: req.RemoveVolumes,
		RemoveLinks:   req.RemoveLinks,
		Force:         req.Force,
	})
}

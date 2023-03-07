package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dockerTypes "github.com/docker/docker/api/types"
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

func (l *RunContainerLogic) RunContainer(req *types.RunContainerRequest) (resp *types.RunContainerResponse, err error) {
	conf := req.Config.DockerFormat()
	_, err = l.svcCtx.DockerClient.ImagePull(l.ctx, req.Config.Image, dockerTypes.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.DockerClient.ContainerCreate(l.ctx, &conf, nil, nil, nil, req.ContainerName)
	if err != nil {
		return nil, err
	}
	resp = new(types.RunContainerResponse)
	resp.ID = res.ID
	return resp, nil
}

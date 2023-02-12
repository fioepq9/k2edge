package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListContainersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContainersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContainersLogic {
	return &ListContainersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContainersLogic) ListContainers(req *types.ListContainersRequest) (resp *types.ListContainersResponse, err error) {
	// todo: add your logic here and delete this line
	dockerContainers, err := l.svcCtx.DockerClient.ContainerList(context.Background(), dockerTypes.ContainerListOptions{
		Size:   req.Size,
		All:    req.All,
		Latest: req.Latest,
		Since:  req.Since,
		Before: req.Before,
		Limit:  req.Limit,
	})
	if err != nil {
		return
	}
	// todo: find metadata use id and node-id
	//

	containers := make([]types.Container, len(dockerContainers))
	for i := range containers {
		containers[i] = types.NewContainer(
			types.ContainerBuildOptions.FromDocker(dockerContainers[i]),
			types.ContainerBuildOptions.WithNamespace("todo"),
		)
	}
	return
}

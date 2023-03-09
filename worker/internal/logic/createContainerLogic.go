package logic

import (
	"context"
	"fmt"
	"io"
	"time"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContainerLogic {
	return &CreateContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContainerLogic) CreateContainer(req *types.CreateContainerRequest) (resp *types.CreateContainerResponse, err error) {
	d := l.svcCtx.DockerClient
	var conf container.Config
	if len(req.Config.Command) != 0 {
		conf.Cmd = []string{req.Config.Command}
		conf.Cmd = append(conf.Cmd, req.Config.Args...)
	}
	conf.Image = req.Config.Image
	conf.Env = req.Config.Env
	rd, err := d.ImagePull(l.ctx, conf.Image, dockerTypes.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	defer rd.Close()
	_, err = io.ReadAll(rd)
	if err != nil {
		return nil, err
	}
	pbs := make(nat.PortMap)
	for _, e := range req.Config.Expose {
		port := nat.Port(fmt.Sprintf("%d/%s", e.Port, e.Protocol))
		pb := []nat.PortBinding{
			{HostPort: fmt.Sprint(e.HostPort)},
		}
		pbs[port] = pb
	}
	res, err := l.svcCtx.DockerClient.ContainerCreate(
		l.ctx,
		&conf,
		&container.HostConfig{
			PortBindings: pbs,
		},
		nil,
		nil,
		fmt.Sprintf("%s-%d", req.ContainerName, time.Now().Unix()),
	)
	if err != nil {
		return nil, err
	}
	err = d.ContainerStart(l.ctx, res.ID, dockerTypes.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}
	resp = new(types.CreateContainerResponse)
	resp.ID = res.ID
	return resp, nil
}

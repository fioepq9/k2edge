package logic

import (
	"context"
	"fmt"
	"io"
	"time"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
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
	d := l.svcCtx.Docker
	// 构建 docker 中的 container config
	conf := container.Config{
		Image: req.Config.Image,
		Env:   req.Config.Env,
	}
	if len(req.Config.Command) != 0 {
		conf.Cmd = append([]string{req.Config.Command}, req.Config.Args...)
	}
	// 构建 docker 中的 container host config
	hostConf := container.HostConfig{
		PortBindings: exposedPortToPortMap(req.Config.Expose),
		Resources: container.Resources{
			NanoCPUs: req.Config.Limit.CPU,
			Memory:   req.Config.Limit.Memory,
		},
	}

	// 拉取镜像
	rd, err := d.ImagePull(l.ctx, conf.Image, dtypes.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	defer rd.Close()
	_, err = io.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	// 创建镜像
	res, err := d.ContainerCreate(
		l.ctx,
		&conf,
		&hostConf,
		nil,
		nil,
		fmt.Sprintf("%s-%d", req.ContainerName, time.Now().Unix()),
	)
	if err != nil {
		return nil, err
	}

	// 启动镜像
	err = d.ContainerStart(l.ctx, res.ID, dtypes.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	// 返回镜像 ID
	resp = new(types.CreateContainerResponse)
	resp.ID = res.ID
	return resp, nil
}

func exposedPortToPortMap(ep []types.ExposedPort) nat.PortMap {
	pbs := make(nat.PortMap)
	for _, e := range ep {
		port := nat.Port(fmt.Sprintf("%d/%s", e.Port, e.Protocol))
		pb := []nat.PortBinding{
			{HostPort: fmt.Sprint(e.HostPort)},
		}
		pbs[port] = pb
	}
	return pbs
}

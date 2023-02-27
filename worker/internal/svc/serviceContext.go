package svc

import (
	"context"
	"k2edge/worker/internal/config"
	"k2edge/worker/internal/middleware"

	"github.com/zeromicro/go-zero/rest"

	"github.com/docker/docker/client"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	DockerClient   *client.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	_, err = dockerCli.Ping(context.TODO())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		DockerClient:   dockerCli,
	}
}

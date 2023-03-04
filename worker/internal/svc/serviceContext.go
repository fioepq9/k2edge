package svc

import (
	"context"
	"k2edge/worker/internal/config"
	"k2edge/worker/internal/middleware"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/rest"

	"github.com/docker/docker/client"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	C    *ServiceContext
	once sync.Once
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	DockerClient   *client.Client
	Etcd           *clientv3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	if C == nil {
		once.Do(func() {
			dockerCli, err := client.NewClientWithOpts(client.FromEnv)
			if err != nil {
				panic(err)
			}
			_, err = dockerCli.Ping(context.TODO())
			if err != nil {
				panic(err)
			}
			config := clientv3.Config{
				Endpoints:   c.Etcd.Endpoints,
				DialTimeout: time.Duration(c.Etcd.DialTimeout) * time.Second,
			}
			etcd, err := clientv3.New(config)
			if err != nil {
				panic(err)
			}
			C = &ServiceContext{
				Config:         c,
				AuthMiddleware: middleware.NewAuthMiddleware().Handle,
				DockerClient:   dockerCli,
				Etcd:           etcd,
			}
		})
	}
	return C
}

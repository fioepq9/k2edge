package svc

import (
	"context"
	"fmt"
	"k2edge/worker/internal/config"
	"k2edge/worker/internal/middleware"
	"os"
	"time"

	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	Docker         *client.Client
	Etcd           *clientv3.Client
	Websocket      websocket.Upgrader
}

func NewServiceContext(c config.Config) *ServiceContext {
	d, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	_, err = d.Ping(context.TODO())
	if err != nil {
		panic(err)
	}
	e, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: time.Duration(c.Etcd.DialTimeout) * time.Second,
	})
	if err != nil {
		panic(err)
	}
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	c.Name, err = os.Hostname()
	if err != nil {
		c.Name = fmt.Sprintf("worker-%d", time.Now().Unix())
	}
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		Docker:         d,
		Etcd:           e,
		Websocket:      u,
	}
}

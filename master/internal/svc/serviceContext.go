package svc

import (
	"context"
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/config"

	"k2edge/master/internal/schedule"
	"k2edge/master/internal/types"
	"time"

	"github.com/gorilla/websocket"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceContext struct {
	Config    config.Config
	Etcd      *clientv3.Client
	Websocket websocket.Upgrader
	Event     chan types.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	config := clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: time.Duration(c.Etcd.DialTimeout) * time.Second,
	}

	event := 

	etcd, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &ServiceContext{
		Config:    c,
		Etcd:      etcd,
		Websocket: u,
	}
}

type WorkerFilter func([]types.Node, *types.Container) ([]types.Node, error)

func (s *ServiceContext) Worker(container *types.Container) (*types.Node, error) {
	nodes, err := etcdutil.GetOne[types.Node](s.Etcd, context.TODO(), "/node/"+etcdutil.SystemNamespace)
	if err != nil {
		return nil, err
	}

	*nodes, err = schedule.Schedule(*nodes, container, s.Etcd)
	if err != nil {
		return nil, err
	}

	if len(*nodes) == 0 {
		return nil, fmt.Errorf("not worker can run")
	}
	return  &(*nodes)[0], nil
}

package svc

import (
	"context"
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/config"
	"k2edge/master/internal/types"
	"time"

	"github.com/gorilla/websocket"
	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceContext struct {
	Config    config.Config
	Etcd      *clientv3.Client
	Websocket websocket.Upgrader
}

func NewServiceContext(c config.Config) *ServiceContext {
	config := clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: time.Duration(c.Etcd.DialTimeout) * time.Second,
	}

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

func (s *ServiceContext) Worker(container *types.Container, filters ...WorkerFilter) (*types.Node, error) {
	nodes, err := etcdutil.GetOne[types.Node](s.Etcd, context.TODO(), "/node/"+etcdutil.SystemNamespace)
	if len(filters) == 0 {
		// 在线调度算法
		for _, filter := range filters {
			*nodes, err = filter(*nodes, container)
			if err != nil {
				return nil, err
			}
		}
	}
	
	if err != nil {
		return nil, err
	}
	workers := lo.Filter(*nodes, func(item types.Node, _ int) bool {
		return lo.Contains(item.Roles, "worker")
	})
	for _, f := range filters {
		workers, err = f(workers, container)
		if err != nil {
			return nil, err
		}
	}
	if len(workers) == 0 {
		return nil, fmt.Errorf("not worker can run")
	}
	return &workers[0], nil
}

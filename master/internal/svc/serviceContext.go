package svc

import (
	"k2edge/master/internal/config"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceContext struct {
	Config config.Config
	Etcd   *clientv3.Client
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
	return &ServiceContext{
		Config: c,
		Etcd:   etcd,
	}
}

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Etcd EtcdConf
}

type EtcdConf struct {
	Endpoints   []string
	DialTimeout int64
}
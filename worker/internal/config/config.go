package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Etcd   EtcdConf
	Secret string
}

type EtcdConf struct {
	Endpoints   []string
	DialTimeout int64
}

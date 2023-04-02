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

type APIServer struct {
	Server string // 使用kubectl需要设置Master节点的地址
}
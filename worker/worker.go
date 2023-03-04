package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"k2edge/worker/internal/config"
	"k2edge/worker/internal/handler"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var configFile = flag.String("f", "etc/worker-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	opts := make([]rest.RunOption, 0)
	if c.Mode == "dev" {
		opts = append(opts, rest.WithCors("*"))
	}
	server := rest.MustNewServer(c.RestConf, opts...)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	close, err := RegisterWorker(ctx)
	if err != nil {
		panic(err)
	}
	defer close()
	server.Start()
}

func RegisterWorker(ctx *svc.ServiceContext) (func() error, error) {
	c, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	gresp, err := ctx.Etcd.KV.Get(c, "workers")
	if err != nil {
		return nil, err
	}
	val := gresp.Kvs[0].Value
	var workers []types.Node
	err = json.Unmarshal(val, &workers)
	if err != nil {
		return nil, err
	}
	for _, w := range workers {
		if w.Metadata.Name == ctx.Config.Name {
			return nil, fmt.Errorf("exist name: %s", ctx.Config.Name)
		}
	}
	node := types.Node{
		Metadata: types.Metadata{
			Namespace: "",
			Kind:      "node",
			Name:      ctx.Config.Name,
		},
		BaseURL:      fmt.Sprintf("%s:%d", ctx.Config.Host, ctx.Config.Port),
		Status:       "",
		RegisterTime: 0,
	}
	workers = append(workers, node)
	b, err := json.Marshal(workers)
	if err != nil {
		return nil, err
	}
	_, err = ctx.Etcd.KV.Put(c, "workers", string(b))
	if err != nil {
		return nil, err
	}
	return func() error {
		c, _ := context.WithTimeout(context.TODO(), 10*time.Second)
		workers, err := Get[[]types.Node](ctx.Etcd, c, "workers")
		if err != nil {
			return err
		}
		newWorkers := make([]types.Node, 0)
		for _, w := range *workers {
			if w.Metadata.Name == ctx.Config.Name {
				continue
			}
			newWorkers = append(newWorkers, w)
		}
		Put()

	}, nil
}

func Get[T any](cli *clientv3.Client, ctx context.Context, key string) (result *T, err error) {
	c, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	gresp, err := cli.KV.Get(c, "workers")
	if err != nil {
		return nil, err
	}
	val := gresp.Kvs[0].Value
	var ret T
	err = json.Unmarshal(val, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

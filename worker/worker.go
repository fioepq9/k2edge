package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"k2edge/etcdutil"
	"k2edge/worker/internal/config"
	"k2edge/worker/internal/handler"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
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
	c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	workers, err := etcdutil.GetOne[[]types.Node](ctx.Etcd, c, "workers")
	if err != nil {
		return nil, err
	}
	for _, w := range *workers {
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
		Status:       "active",
		RegisterTime: time.Now().Unix(),
	}
	*workers = append(*workers, node)
	etcdutil.PutOne(ctx.Etcd, c, "workers", *workers)
	return func() error {
		c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		workers, err := etcdutil.GetOne[[]types.Node](ctx.Etcd, c, "workers")
		if err != nil {
			return err
		}
		newWorkers := make([]types.Node, 0)
		for _, w := range *workers {
			if w.Metadata.Name == ctx.Config.Name {
				w.Status = "disconnect"
			}
			newWorkers = append(newWorkers, w)
		}
		return etcdutil.PutOne(ctx.Etcd, c, "workers", newWorkers)
	}, nil
}

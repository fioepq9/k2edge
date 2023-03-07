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

	"github.com/samber/lo"
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
	err := doRegisterWorker(ctx)
	if err != nil {
		return nil, err
	}
	return func() error {
		c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		workersPtr, err := etcdutil.GetOne[[]types.Node](ctx.Etcd, c, "workers")
		if err != nil {
			return err
		}
		workers := *workersPtr
		lo.ForEach(workers, func(_ types.Node, i int) {
			if workers[i].Metadata.Name == ctx.Config.Name {
				workers[i].Status = "closed"
			}
		})
		return etcdutil.PutOne(ctx.Etcd, c, "workers", workers)
	}, nil
}

func doRegisterWorker(ctx *svc.ServiceContext) error {
	c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	workersPtr, err := etcdutil.GetOne[[]types.Node](ctx.Etcd, c, "workers")
	if err != nil {
		return err
	}
	workers := *workersPtr
	item, idx, found := lo.FindIndexOf(workers, func(item types.Node) bool {
		return item.Metadata.Name == ctx.Config.Name
	})
	if found {
		if item.Status == "active" {
			return fmt.Errorf("exist name: %s", ctx.Config.Name)
		} else {
			workers[idx].Status = "active"
			etcdutil.PutOne(ctx.Etcd, c, "workers", workers)
		}
		return nil
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
	workers = append(workers, node)
	etcdutil.PutOne(ctx.Etcd, c, "workers", workers)
	return nil
}

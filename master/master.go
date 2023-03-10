package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/config"
	"k2edge/master/internal/handler"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/master-api.yaml", "the config file")
var registerNamespace = "" 

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
	close, err := RegisterMaster(ctx)
	if err != nil {
		panic(err)
	}
	defer close()
	go server.Start()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	for {
		sig := <-ch
		if sig == os.Interrupt {
			return
		}
	}
}

func RegisterMaster(ctx *svc.ServiceContext) (func() error, error) {
	err := doRegisterMaster(ctx)
	if err != nil {
		return nil, err
	}
	return func() error {
		c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		nodesPtr, err := etcdutil.GetOne[[]types.Node](ctx.Etcd, c, "/nodes")
		if err != nil {
			return err
		}
		nodes := *nodesPtr
		lo.ForEach(nodes, func(_ types.Node, i int) {
			if nodes[i].Metadata.Name == ctx.Config.Name {
				nodes[i].Status = "closed"
			}
		})
		return etcdutil.PutOne(ctx.Etcd, c, "/nodes", nodes)
	}, nil
}

func doRegisterMaster(ctx *svc.ServiceContext) error {
	c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	nodesPtr, err := etcdutil.GetOne[[]types.Node](ctx.Etcd, c, "/nodes")
	if err != nil && !errors.Is(err, etcdutil.ErrKeyNotExist) {
		return err
	}
	nodes := make([]types.Node, 0)

	if nodesPtr != nil {
		nodes = *nodesPtr
		item, idx, found := lo.FindIndexOf(nodes, func(item types.Node) bool {
			return item.Metadata.Name == ctx.Config.Name && item.Metadata.Namespace == registerNamespace
		})
		
		if found {
			if item.Status == "active" {
				if lo.Contains(item.Roles, "master") {
					return fmt.Errorf("master already exist")
				}
				nodes[idx].Roles = append(nodes[idx].Roles, "master")
			} else {
				nodes[idx].Status = "active"
				nodes[idx].Roles = []string{"master"}
			}
			etcdutil.PutOne(ctx.Etcd, c, "/nodes", nodes)
			return nil
		}

	}

	node := types.Node{
		Metadata: types.Metadata{
			Namespace: registerNamespace,
			Kind:      "node",
			Name:      ctx.Config.Name,
		},
		Roles:        []string{"master"},
		BaseURL:      fmt.Sprintf("http://%s:%d", ctx.Config.Host, ctx.Config.Port),
		Status:       "active",
		RegisterTime: time.Now().Unix(),
	}
	nodes = append(nodes, node)
	etcdutil.PutOne(ctx.Etcd, c, "/nodes", nodes)
	return nil
}
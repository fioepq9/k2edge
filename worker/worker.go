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
	"k2edge/worker/event"
	"k2edge/worker/internal/config"
	"k2edge/worker/internal/handler"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/worker-api.yaml", "the config file")
var registerNamespace = etcdutil.SystemNamespace

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

	go server.Start()

	// 启动 docker 事件监听
	e := event.NewEventSubcriber(ctx)
	go e.Subcribe()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	for {
		sig := <-ch
		if sig == os.Interrupt {
			return
		}
	}
}

func RegisterWorker(ctx *svc.ServiceContext) (func() error, error) {
	err := doRegisterWorker(ctx)
	if err != nil {
		return nil, err
	}
	return func() error {
		c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()
		key := "/node/" + registerNamespace + "/" + ctx.Config.Name
		nodesPtr, err := etcdutil.GetOne[types.Node](ctx.Etcd, c, key)
		if err != nil {
			return err
		}
		nodes := *nodesPtr
		lo.ForEach(nodes, func(_ types.Node, i int) {
			if nodes[i].Metadata.Name == ctx.Config.Name {
				nodes[i].Status.Working = false
			}
		})
		return etcdutil.PutOne(ctx.Etcd, c, key, nodes[0])
	}, nil
}

func doRegisterWorker(ctx *svc.ServiceContext) error {
	key := "/node/" + registerNamespace + "/" + ctx.Config.Name
	c, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	node, err := etcdutil.GetOne[types.Node](ctx.Etcd, c, key)
	if err != nil && !errors.Is(err, etcdutil.ErrKeyNotExist) {
		return err
	}
	var n types.Node
	if node != nil {
		n := (*node)[0]
		if n.Status.Working {
			// 结点原本被注册为 worker
			if lo.Contains(n.Roles, "worker") {
				return fmt.Errorf("worker already exist")
			}

			// 结点原本被注册为 master
			n.Roles = append(n.Roles, "worker")
			n.BaseURL.WorkerURL = fmt.Sprintf("http://%s:%d", ctx.Config.Host, ctx.Config.Port)
		} else {
			n.Status.Working = true
			n.Roles = []string{"worker"}
		}
		etcdutil.PutOne(ctx.Etcd, c, key, n)
		return nil
	}

	// 结点原本没被注册过
	n = types.Node{
		Metadata: types.Metadata{
			Namespace: registerNamespace,
			Kind:      "node",
			Name:      ctx.Config.Name,
		},
		Roles: []string{"worker"},
		BaseURL: types.NodeURL{
			WorkerURL: fmt.Sprintf("http://%s:%d", ctx.Config.Host, ctx.Config.Port),
			MasterURL: "",
		},
		Spec: types.Spec{
			Unschedulable: false,
		},
		RegisterTime: time.Now().Unix(),
		Status: types.Status{
			Working: true,
		},
	}
	etcdutil.PutOne(ctx.Etcd, c, "/node/"+registerNamespace+"/"+n.Metadata.Name, n)
	return nil
}

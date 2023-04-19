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
	"k2edge/master/monitor"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/master-api.yaml", "the config file")
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
	close, err := RegisterMaster(ctx)
	if err != nil {
		panic(err)
	}
	defer close()
	go server.Start()
	go monitor.EventMonitor(ctx)
	go monitor.StatusMonitor(ctx)
	go monitor.CornjobMonitor(ctx)
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

func doRegisterMaster(ctx *svc.ServiceContext) error {
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
			// 结点原本被注册为 master
			if lo.Contains(n.Roles, "master") {
				return fmt.Errorf("master already exist")
			}

			// 结点原本被注册为 worker
			n.Roles = append(n.Roles, "master")
			n.BaseURL.MasterURL = fmt.Sprintf("http://%s:%d", ctx.Config.Host, ctx.Config.Port)
		} else {
			n.Status.Working = true
			n.Roles = []string{"master"}
		}
		etcdutil.PutOne(ctx.Etcd, c, key, n)
		return nil
	}

	// 结点原本没被注册过
	// CPU
	cpuCount, err := cpu.CountsWithContext(context.Background(), true)
	if err != nil {
		panic(err)
	}
	cpuTotal := float64(cpuCount) * 1e9
	// Memory
	memStat, err := mem.VirtualMemoryWithContext(context.Background())
	if err != nil {
		panic(err)
	}
	memoryTotal := memStat.Total
	n = types.Node{
		Metadata: types.Metadata{
			Namespace: registerNamespace,
			Kind:      "node",
			Name:      ctx.Config.Name,
		},
		Roles: []string{"master"},
		BaseURL: types.NodeURL{
			WorkerURL: "",
			MasterURL: fmt.Sprintf("http://%s:%d", ctx.Config.Host, ctx.Config.Port),
		},
		Spec: types.Spec{
			Unschedulable: false,
			Capacity: types.Capacity{
				CPU:    int64(cpuTotal),
				Memory: int64(memoryTotal),
			},
		},
		RegisterTime: time.Now().Unix(),
		Status: types.Status{
			Working: true,
		},
	}
	etcdutil.PutOne(ctx.Etcd, c, "/node/"+etcdutil.SystemNamespace+"/"+n.Metadata.Name, n)
	return nil
}

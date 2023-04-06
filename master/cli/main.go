package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/types"
	cmd "k2edge/master/cli/command"

	"github.com/samber/lo"
	"github.com/urfave/cli"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
)

func main() {
	var filePath string = "../etc/master-api.yaml"

	app := cli.NewApp()
	app.Before = func(ctx *cli.Context) error {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("k2e get configuration failed")
		}
		
		
		config := new(config)
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return fmt.Errorf("k2e get configuration failed")
		}

		if len(config.Etcd.Endpoints) == 0 {
			panic("k2e initial: cannot found etcd")
		} 
		
		server := getServer(context.Background(), config.Etcd.Endpoints[0])
		// 将一些全局配置信息添加到上下文中
		ctx.App.Metadata = map[string]interface{}{
			"config": map[string]string{
                "server": server,
            },
		}
		return nil
	}

	app.Commands = []cli.Command{*cmd.Namespace() }

	err := app.Run(os.Args)
    if err != nil {
        fmt.Println(err)
    }
}

func getServer(ctx context.Context, endPoint string) string {
	config := clientv3.Config{
		Endpoints:   []string{endPoint},
		DialTimeout: 5 * time.Second,
	}

	etcd, err := clientv3.New(config)
	if err != nil {
		panic(fmt.Errorf("k2e initial: %s", err))
	}
	// 获取node信息
	nodes, err := etcdutil.GetOne[types.Node](etcd, ctx, "/node/" + etcdutil.SystemNamespace)

	if err != nil {
		panic(err)
	}

	for _, node := range *nodes {
		if node.Status.Working && lo.Contains(node.Roles, "master") {
			return node.BaseURL.MasterURL
		}
	}
	panic("k2e initial: cannot found master node")
}

type config struct {
	Etcd Etcd `yaml:"Etcd"`
}

type Etcd struct {	
	Endpoints []string `yaml:"Endpoints"`
	DialTimeout int		`yaml:"DialTimeout"`
}
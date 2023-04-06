package cli

import (
	"context"
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/types"
	"time"

	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"
)


func getServer(endPoint string) string {
	config := clientv3.Config{
		Endpoints:   []string{endPoint},
		DialTimeout: 5 * time.Second,
	}

	etcd, err := clientv3.New(config)
	if err != nil {
		panic(fmt.Errorf("k2e initial: %s", err))
	}
	// 获取node信息
	nodes, err := etcdutil.GetOne[types.Node](etcd, context.Background(), "/node/" + etcdutil.SystemNamespace)

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
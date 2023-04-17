package schedule

import (
	"context"
	"fmt"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"
)


type nodeInfo struct {
	etcdInfo types.Node
	actualInfo   types.NodeTopResponse
	score  float64
}

type Scheduler struct {
	nodeInfo     []nodeInfo
	container *types.Container
	etcd *clientv3.Client
	Err error
}

func NewScheduler(nodes []types.Node, container *types.Container, etcd *clientv3.Client) *Scheduler {
    nodes, err := PreProcessing(nodes)

	if err != nil {
		return &Scheduler{
			Err: err,
		}
	}

	s := &Scheduler{
		container: container,
		nodeInfo: make([]nodeInfo, 0),
		etcd: etcd,
	}

	//获取node的实际信息
	for _, node := range nodes {
		var n nodeInfo
		cli := client.NewClient(node.BaseURL.WorkerURL)
		topInfo, err := cli.Node.Top(context.TODO())
		if err != nil {
			s.Err = err
			return s
		}
		n.etcdInfo = node
		n.actualInfo = types.NodeTopResponse(*topInfo)
		n.score = 0
		s.nodeInfo = append(s.nodeInfo, n)
	}

	return s
}

func (s *Scheduler) GetNodes() ([]types.Node, error) {
	nodes := lo.Map(s.nodeInfo, func (item nodeInfo, _ int) types.Node {
		return item.etcdInfo
	})
	return nodes, s.Err
}

// 除去master结点和不可调度结点
func PreProcessing(nodes []types.Node) ([]types.Node, error) {
	nodes = lo.Filter(nodes, func(item types.Node, _ int) bool {
		return lo.Contains(item.Roles, "worker") && item.Status.Working && !item.Spec.Unschedulable
	})

	return nodes, nil
}


func Schedule(nodes []types.Node, container *types.Container, etcd *clientv3.Client) ([]types.Node, error) {
	s := NewScheduler(nodes, container, etcd).Predicate().Priority()
	for _,i := range s.nodeInfo {
		fmt.Printf("%s %f\n", i.etcdInfo.Metadata.Name, i.score)
	}
	return s.GetNodes()
}

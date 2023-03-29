package schedule

import (
	"context"
	"fmt"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/samber/lo"
)


type nodeInfo struct {
	config types.Node
	info   types.NodeTopResponse
	score  float64
}

type Scheduler struct {
	nodeInfo     []nodeInfo
	container *types.Container
	Err error
}

func NewScheduler(nodes []types.Node, container *types.Container) *Scheduler {
    nodes, err := PreProcessing(nodes)

	if err != nil {
		return &Scheduler{
			Err: err,
		}
	}

	s := &Scheduler{
		container: container,
		nodeInfo: make([]nodeInfo, 0),
	}

	for _, node := range nodes {
		var n nodeInfo
		cli := client.NewClient(node.BaseURL.WorkerURL)
		topInfo, err := cli.Node.Top(context.TODO())
		if err != nil {
			s.Err = err
			return s
		}
		n.config = node
		n.info = types.NodeTopResponse(*topInfo)
		n.score = 0
		s.nodeInfo = append(s.nodeInfo, n)
	}

	fmt.Print(s.nodeInfo)
	fmt.Println("NewScheduler")

	return s
}

func (s *Scheduler) GetNodes() ([]types.Node, error) {
	fmt.Println(s.nodeInfo)
	nodes := lo.Map(s.nodeInfo, func (item nodeInfo, _ int) types.Node {
		return item.config
	})
	return nodes, s.Err
}


// 除去workering结点和不可调度结点
func PreProcessing(nodes []types.Node) ([]types.Node, error) {
	nodes = lo.Filter(nodes, func(item types.Node, _ int) bool {
		return lo.Contains(item.Roles, "worker") && item.Status.Working && !item.Spec.Unschedulable
	})

	fmt.Print(nodes)
	fmt.Println("PreProcessing")
	return nodes, nil
}


func Schedule(nodes []types.Node, container *types.Container) ([]types.Node, error) {
	return NewScheduler(nodes, container).Predicate().Priority().GetNodes()
}

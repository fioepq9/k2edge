package schedule

import (
	"context"
	"k2edge/master/internal/client"
	"k2edge/master/internal/types"

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
	nodes		[]types.Node
	Err error
}

func NewScheduler(nodes []types.Node, container *types.Container) *Scheduler {
	s := &Scheduler{
		nodes:     nodes,
		container: container,
		nodeInfo: make([]nodeInfo, 0),
	}

	nodes = lo.Filter(nodes, func(item types.Node, _ int) bool {
		return lo.Contains(item.Roles, "worker")
	})

	for _, node := range nodes {
		var n nodeInfo
		cli := client.NewClient(node.BaseURL.WorkerURL)
		topInfo, err := cli.Node.Top(context.TODO())
		if err != nil {
			s.Err = err
			return s
		}
		n.config = node
		n.info = *topInfo
		n.score = 0
		s.nodeInfo = append(s.nodeInfo, n)
	}
	return s
}

func (s *Scheduler) GetNodes() ([]types.Node, error) {
	return s.nodes, s.Err
}


func Schedule(nodes []types.Node, container *types.Container) ([]types.Node, error) {
	return NewScheduler(nodes, container).Predicate().Priority().GetNodes()
}

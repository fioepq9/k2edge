package schedule

import (
	"context"
	"k2edge/master/internal/client"
	"k2edge/master/internal/types"

	"github.com/samber/lo"
)

type nodeInfo struct {
	config types.Node
	info types.NodeTopResponse
}

func Scheduler(nodes []types.Node, container *types.Container) ([]types.Node, error) {
	info := make([]nodeInfo, 0)
	for _, node := range nodes {
		var n nodeInfo
		cli := client.NewClient(node.BaseURL.MasterURL)
		topInfo, err := cli.Node.Top(context.TODO())
		if err != nil {
			return nil, err
		}
		
		n.config = node
		n.info = *topInfo
		info = append(info, n)
	}
	
	info, err := predicate(info, container)
	if err != nil {
		return nil, err
	}

	set := make(map[string]bool)
	for _, i := range info  {
		set[i.config.Metadata.Name] = true
	}

	nodes = lo.Filter(nodes, func (item types.Node, idx int) bool {
		_, found := set[item.Metadata.Name]
		return found
	})

	return nodes, nil
}


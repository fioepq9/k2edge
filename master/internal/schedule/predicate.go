package schedule

import (
	"k2edge/master/internal/types"

	"github.com/samber/lo"
)



// 过滤算法，过滤掉不符合要求的node
func predicate(nodes []nodeInfo, container *types.Container) ([]nodeInfo, error) {
	nodes, err := podFitsHost(nodes, container)
	if err != nil {
		return nil, err
	}
	
	nodes, err = podFitHostPorts(nodes, container)
	if err != nil {
		return nil, err
	}

	nodes, err = podFitsResource(nodes, container)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// 检查容器是否指定了某个node
func podFitsHost(nodes []nodeInfo, container *types.Container) ([]nodeInfo, error) {
	return lo.Filter(nodes, func(item nodeInfo, index int) bool {
		return item.config.Metadata.Name == container.ContainerConfig.NodeName
	}), nil
}

// 检查pod需要的端口，在结点上是否可用
func podFitHostPorts(nodes []nodeInfo, container *types.Container) ([]nodeInfo, error) {
	//待实现
	return nodes, nil
}

// 检查node是否有空闲资源以满足容器需求
func podFitsResource(nodes []nodeInfo, container *types.Container) ([]nodeInfo, error) {
	lo.Filter(nodes, func(item nodeInfo, index int) bool {
		return (item.info.CPUFree * 0.95) > float64(container.ContainerConfig.Request.CPU) && 
				(float64(item.info.MemoryAvailable) * 0.95) > float64(container.ContainerConfig.Request.Memory)
	})
	return nodes, nil
}

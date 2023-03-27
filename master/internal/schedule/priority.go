package schedule

import (
	"k2edge/master/internal/types"
	"sort"

	"github.com/samber/lo"
)

// 打分算法，为过滤出的node进行打分
func (s *Scheduler) Priority() *Scheduler {
	if s.Err != nil {
		return s
	}

	return s.SelectorSpreadPriority() .SortPriority()
}

// 将node按nodeInfo中的分数排序
func (s *Scheduler) SortPriority() *Scheduler {
	
	sort.Slice(s.nodeInfo, func(i, j int) bool {
		return s.nodeInfo[i].score < s.nodeInfo[j].score
	})

	s.nodes = lo.Map(s.nodeInfo, func(item nodeInfo, _ int) types.Node {
		return item.config
	})
	return s
}

// 优先减少node上属于同一个Deployment的pod数量
// • 首先获取正在调度的容器有关的Deployment
// • 对于每一个node来说：
//   - 获取该node上的所有Deployment，在所有的Deployment中，如果有与正在调度的容器的Deployment相同的，node的分数+1
//   - 所有node计算完分数之后，选取node的最大值，进行分数调整
//   - 遍历所有node值 ，计算node的分数为  10✖(最大值-当前node值)/最大值
func (s *Scheduler) SelectorSpreadPriority() *Scheduler {
	if s.Err != nil {
		return s
	}

	const ratio = 1.0

	// 获取正在调度的容器有关的Deployment

	// 从获取的Deployment中，得到Deployment有关的容器和其运行的node，这些node分数+1

	// 获取node分数最大值

	// 调整node的得分,包含该部分的权重

	return s
}

// 节点上放置的容器越多，这些容器使用的资源越多，这个Node给出的打分就越低，所以优先调度到容器资源占比少的node上。
// 该函数会计算每个node的资源利用率，它目前只考虑CPU和内存两种资源，而且的资源公式为：
// • node分数=（CPU资源可利用率✖10 + 内存资源可利用率✖10）/ 2
// • 资源剩余可利用率  = （资源容量  - 容器要求的资源）/ 资源容量
// 容器要求资源越大，得分越低
func LeastRequestedPriority(nodes []nodeInfo, container *types.Container, ratio float64) ([]nodeInfo, error) {
	for _, node := range nodes {
		score := (1 - float64(container.ContainerConfig.Request.CPU)/node.info.CPUTotal) * 10
		score += (1 - float64(container.ContainerConfig.Request.Memory)/float64(node.info.MemoryTotal)) * 10
		score /= 2

		//权重调整
		score *= ratio
	}
	return nodes, nil
}

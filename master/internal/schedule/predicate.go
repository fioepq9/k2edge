package schedule

import (
	"github.com/samber/lo"
)

// 过滤算法，过滤掉不符合要求的node
func (s *Scheduler) Predicate() *Scheduler {
	if s.Err != nil {
		return s
	}

	return s.PodFitsHost().//PrintScore("指定节点").
			 PodFitHostPorts().//PrintScore("检查端口").
			 PodFitsResource()//PrintScore("节点压力")
}

// 检查容器是否指定了某个node
func (s *Scheduler) PodFitsHost() *Scheduler {
	if s.Err != nil {
		return s
	}

	if s.container.ContainerConfig.NodeName != "" {
		s.nodeInfo = lo.Filter(s.nodeInfo , func(item nodeInfo, index int) bool {
			return item.etcdInfo.Metadata.Name == s.container.ContainerConfig.NodeName
		})
	}
	
	return s
}

// 检查pod需要的端口，在结点上是否可用
func (s *Scheduler)  PodFitHostPorts() *Scheduler {
	if s.Err != nil {
		return s
	}

	//端口扫描
	return s
}

// 检查node是否有空闲资源以满足容器需求
func (s *Scheduler)  PodFitsResource() *Scheduler {
	if s.Err != nil {
		return s
	}

	s.nodeInfo = lo.Filter(s.nodeInfo, func(item nodeInfo, index int) bool {
		return (item.actualInfo.CPUFree * 0.95) > float64(s.container.ContainerConfig.Request.CPU) && 
				(float64(item.actualInfo.MemoryAvailable) * 0.95) > float64(s.container.ContainerConfig.Request.Memory)
	})
	return s
}

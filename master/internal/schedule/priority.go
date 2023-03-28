package schedule

import (
	"context"
	"k2edge/master/internal/types"
	"math"
	"sort"

	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

// 打分算法，为过滤出的node进行打分，每个打分算法出来的分数范围在 0 ～ 10
func (s *Scheduler) Priority() *Scheduler {
	if s.Err != nil {
		return s
	}

	return s.SelectorSpreadPriority().LeastRequestedPriority().BalancedResourceAllocation().ImageLocalityPriority().MemoryPressure().CPUPressure().SortPriority()
}

// 将node按nodeInfo中的分数排序
func (s *Scheduler) SortPriority() *Scheduler {
	if s.Err != nil {
		return s
	}

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

	//const ratio = 1.0

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
func (s *Scheduler) LeastRequestedPriority()  *Scheduler {
	if s.Err != nil {
		return s
	}

	// 权重
	const ratio = 1.0

	for idx, info := range s.nodeInfo {
		score := (1 - float64(s.container.ContainerConfig.Request.CPU)/info.info.CPUTotal) * 10
		score += (1 - float64(s.container.ContainerConfig.Request.Memory)/float64(info.info.MemoryTotal)) * 10
		score /= 2

		//权重调整
		score *= ratio

		s.nodeInfo[idx].score += score
	}
	return s
}


// 尽量选择各项资源（CPU，内存）使用率最均衡的node，计算公式为：
// • 1-ABS（（容器请求的CPU资源+node已使用的CPU资源）/node的CPU容量  - 
//     （容器请求的内存资源+node已使用的内存资源/节点的内存容量）
func (s *Scheduler) BalancedResourceAllocation()  *Scheduler {
	if s.Err != nil {
		return s
	}

	// 权重
	const ratio = 1.0

	for idx, info := range s.nodeInfo {
		score := 10 - math.Abs((float64(s.container.ContainerConfig.Request.CPU) + info.info.CPUUsed) / info.info.CPUTotal -
	 				(float64(s.container.ContainerConfig.Request.Memory) + float64(info.info.MemoryUsed)) / float64(info.info.MemoryTotal)) * 10

		//权重调整
		score *= ratio
		s.nodeInfo[idx].score += score
	}
	return s
}


// 根据当地node本地是否存在容器所需的镜像计算的得分
// • 计算已经下载容器镜像的node的数量占k2edge总结点数量的比重（叫传播度，防止node heating），
// 	  然后将比重与指定的容器镜像大小相乘，得到sumScores。
// • 如果sumScores < 23 * 1024 * 1024，则得分为0
// • 如果sumScores >= 1000 * 1024 * 1024，则得分为10
// • 包含容器镜像的node，得分为sumScores，不包含容器镜像的node，得分为0
func (s *Scheduler) ImageLocalityPriority()  *Scheduler {
	if s.Err != nil {
		return s
	}

	// 权重
	const ratio = 1.0

	// 计算包含container镜像的node数量
	const MB = 1024 * 1024
	containImageNodeNum := 0
	nodeSet := make(map[string]bool) //记录包含image的node名字
	for _, info := range s.nodeInfo {
		if lo.Contains(info.info.Images, s.container.ContainerConfig.Image) {
			containImageNodeNum++
			nodeSet[info.config.Metadata.Name] = true
		}
	}

	// 获取container镜像大小
	cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        s.Err = err
		return s
    }
    image, _, err := cli.ImageInspectWithRaw(context.Background(), "your-image-name:tag")
    if err != nil {
        s.Err = err
		return s
    }
	imageSize := image.Size

	// 计算最终分数
	sumScore := (float64(containImageNodeNum) / float64(len(s.nodeInfo))) * float64(imageSize)
	if sumScore < 23 * MB {
		sumScore = 0 // 0分
	} else if sumScore > 1000 * MB {
		sumScore = 977 * MB // 10分
	}
	sumScore /= 977 * MB * 10

	// 调整权重
	sumScore *= ratio

	for	idx, info := range s.nodeInfo {
		if _, found := nodeSet[info.config.Metadata.Name]; found {
			s.nodeInfo[idx].score += sumScore
		}
	}

	return s
}

// MemoryPressure 压力影响，达到特定压力值后，这部分分数为0
func (s *Scheduler) MemoryPressure() *Scheduler {
	if s.Err != nil {
		return s
	}

	// 权重
	const ratio = 1.0

	// 调整权重
	score := 10 * ratio

	for idx, info := range s.nodeInfo {
		if info.info.MemoryUsedPercent > 0.8 {
			s.nodeInfo[idx].score += score
		} 
	}

	return s
}

// CPUPressure 压力影响，达到特定压力值后，这部分分数为0
func (s *Scheduler) CPUPressure() *Scheduler {
	if s.Err != nil {
		return s
	}

	// 权重
	const ratio = 1.0

	// 调整权重
	score := 10 * ratio

	for idx, info := range s.nodeInfo {
		if info.info.CPUUsedPercent > 0.8 {
			s.nodeInfo[idx].score += score
		} 
	}

	return s
}

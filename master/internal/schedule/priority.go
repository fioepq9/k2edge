package schedule

import (
	"context"
	"errors"
	"k2edge/etcdutil"
	"k2edge/master/internal/types"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 打分算法，为过滤出的node进行打分，每个打分算法出来的分数范围在 0 ～ 10
func (s *Scheduler) Priority() *Scheduler {
	if s.Err != nil {
		return s
	}

	return s.SelectorSpreadPriority().//PrintScore("deployment").
			LeastRequestedPriority().//PrintScore("最多空闲资源").
			BalancedResourceAllocation().//PrintScore("资源均衡").
			ImageLocalityPriority().//PrintScore("镜像大小").
			MemoryPressure().//PrintScore("内存压力").
			CPUPressure().//PrintScore("CPU压力").
			SortPriority()//.PrintScore("排序")
}

// 将node按nodeInfo中的分数排序
func (s *Scheduler) SortPriority() *Scheduler {
	if s.Err != nil {
		return s
	}

	sort.Slice(s.nodeInfo, func(i, j int) bool {
		return s.nodeInfo[i].score > s.nodeInfo[j].score
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
	if s.Err != nil || s.container.ContainerConfig.Deployment == "" {
		return s
	}

	const ratio = 1.0

	// 获取正在调度的容器有关的Deployment
	nn := strings.Split(s.container.ContainerConfig.Deployment, "/")
	dnamespace := nn[0]
	dname := nn[1]
	deployments, err := etcdutil.GetOne[types.Deployment](s.etcd, context.Background(), etcdutil.GenerateKey("deployment", dnamespace, dname))
	if err != nil {
		if errors.Is(err, etcdutil.ErrKeyNotExist) {
			return s
		}
		s.Err = err
		return s
	}

	// 从获取的Deployment中，得到Deployment有关的容器和其运行的node，这些node分数+1
	deployment := (*deployments)[0]
	set := make(map[string]int)
	max := 0

	for _, c := range deployment.Status.Containers {
		set[c.Node] += 1
		if max < set[c.Node] {
			max = set[c.Node]
		}
	}

	if max == 0 {
		return s
	}

	
	for idx, info := range s.nodeInfo {
		s.nodeInfo[idx].score += (1 - float64(set[info.etcdInfo.Metadata.Name])/float64(max)) * 10 * ratio
	}

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
		score := (1 - float64(s.container.ContainerConfig.Request.CPU)/float64(info.etcdInfo.Spec.Capacity.CPU)) * 5
		score += (1 - float64(s.container.ContainerConfig.Request.Memory)/float64(info.etcdInfo.Spec.Capacity.Memory)) * 5

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
		score := 10 - math.Abs((float64(s.container.ContainerConfig.Request.CPU) + float64(info.etcdInfo.Status.Allocatable.CPU)) / float64(info.etcdInfo.Spec.Capacity.CPU) -
	 				(float64(s.container.ContainerConfig.Request.Memory) + float64(info.etcdInfo.Status.Allocatable.Memory)) / float64(info.etcdInfo.Spec.Capacity.Memory)) * 10

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
	imageSize := 0
	nodeSet := make(map[string]bool) //记录包含image的node名字
	for _, info := range s.nodeInfo {
		for _, i := range info.actualInfo.Images {
			is := strings.Split(i, " ")
			tag := strings.Split(is[0], ":")[0]
			size := is[1]

			var err error
			if tag == s.container.ContainerConfig.Image {
				containImageNodeNum++
				nodeSet[info.etcdInfo.Metadata.Name] = true
				imageSize, err = strconv.Atoi(size)
				if err != nil {
					s.Err = err
					return s
				}
			}
		}
	}

	if containImageNodeNum == 0 {
		return s
	}

	// 计算最终分数
	var sumScore float64
	if imageSize >= 1000 * MB {
		sumScore = (float64(containImageNodeNum) / float64(len(s.nodeInfo))) * 10
	} else if (imageSize < 23) {
		sumScore = 0
	} else {
		sumScore = (float64(containImageNodeNum) / float64(len(s.nodeInfo))) * float64(imageSize) / (100 * MB)
	}

	// 调整权重
	sumScore *= ratio

	for	idx, info := range s.nodeInfo {
		if _, found := nodeSet[info.etcdInfo.Metadata.Name]; found {
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

	for idx, info := range s.nodeInfo {
		percent := float64(info.etcdInfo.Status.Allocatable.Memory) / float64(info.etcdInfo.Spec.Capacity.Memory)
		if percent < 0.8 {
			s.nodeInfo[idx].score += 10 * ratio
		} else if percent >= 0.9 {
			s.nodeInfo[idx].score += 0
		} else {
			s.nodeInfo[idx].score = (90 - 100 * percent) * ratio
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

	for idx, info := range s.nodeInfo {
		percent := float64(info.etcdInfo.Status.Allocatable.CPU) / float64(info.etcdInfo.Spec.Capacity.CPU)
		if percent < 0.8 {
			s.nodeInfo[idx].score += 10 * ratio
		} else if percent >= 0.9 {
			s.nodeInfo[idx].score += 0
		} else {
			s.nodeInfo[idx].score = (90 - 100 * percent) * ratio
		}
	}

	return s
}

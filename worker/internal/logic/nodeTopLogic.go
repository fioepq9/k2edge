package logic

import (
	"context"
	"time"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
	"github.com/samber/lo"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zeromicro/go-zero/core/logx"
)

type NodeTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NodeTopLogic {
	return &NodeTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeTopLogic) NodeTop() (resp *types.NodeTopResponse, err error) {
	d := l.svcCtx.Docker
	resp = new(types.NodeTopResponse)
	// Images
	imagesSumary, err := d.ImageList(l.ctx, dtypes.ImageListOptions{All: true})
	if err != nil {
		return nil, err
	}
	resp.Images = lo.FlatMap(imagesSumary, func(img dtypes.ImageSummary, _ int) []string {
		return img.RepoTags
	})
	// CPU
	cpuInfoStat, err := cpu.InfoWithContext(l.ctx)
	if err != nil {
		return nil, err
	}
	cpuPercent, err := cpu.PercentWithContext(l.ctx, time.Millisecond, true)
	if err != nil {
		return nil, err
	}
	resp.CPU = make([]types.CPUInfo, 0)
	for i := range cpuInfoStat {
		var p float64 = 0
		if i < len(cpuPercent) {
			p = cpuPercent[i]
		}
		resp.CPU = append(resp.CPU, types.CPUInfo{
			CPU:       cpuInfoStat[i].CPU,
			Cores:     cpuInfoStat[i].Cores,
			Mhz:       cpuInfoStat[i].Mhz,
			ModelName: cpuInfoStat[i].ModelName,
			Percent:   p,
		})
	}
	// Memory
	memStat, err := mem.VirtualMemoryWithContext(l.ctx)
	if err != nil {
		return nil, err
	}
	resp.MemoryAvailable = memStat.Available
	resp.MemoryUsed = memStat.Used
	resp.MemoryUsedPercent = memStat.UsedPercent
	resp.MemoryTotal = memStat.Total
	// Disk
	diskStat, err := disk.UsageWithContext(l.ctx, "/")
	if err != nil {
		return nil, err
	}
	resp.DiskFree = diskStat.Free
	resp.DiskUsed = diskStat.Used
	resp.DiskUsedPercent = diskStat.UsedPercent
	resp.DiskTotal = diskStat.Total
	return
}

package logic

import (
	"context"
	"time"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zeromicro/go-zero/core/logx"
)

type HostTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHostTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostTopLogic {
	return &HostTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HostTopLogic) HostTop() (resp *types.NodeTopResponse, err error) {
	resp = new(types.NodeTopResponse)
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

	return resp, nil
}

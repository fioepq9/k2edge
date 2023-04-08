package logic

import (
	"context"
	"errors"
	"time"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	dtypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/samber/lo"
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
	d, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	// Images
	imagesSumary, err := d.ImageList(l.ctx, dtypes.ImageListOptions{All: true})
	if err != nil {
		return nil, err
	}
	resp.Images = lo.FlatMap(imagesSumary, func(img dtypes.ImageSummary, _ int) []string {
		return img.RepoTags
	})
	// CPU
	cpuTimeSet, err := cpu.TimesWithContext(l.ctx, false)
	if err != nil {
		return nil, err
	}
	if len(cpuTimeSet) != 1 {
		return nil, errors.New("get cpu times failed")
	}
	cpuPercent, err := cpu.PercentWithContext(l.ctx, time.Second, false)
	if err != nil {
		return nil, err
	}
	if len(cpuPercent) != 1 {
		return nil, errors.New("get cpu percent failed")
	}
	resp.CPUUsedPercent = cpuPercent[0]
	resp.CPUFree = cpuTimeSet[0].Idle + cpuTimeSet[0].Iowait
	resp.CPUUsed = cpuTimeSet[0].System + cpuTimeSet[0].Nice + cpuTimeSet[0].User + cpuTimeSet[0].Irq + cpuTimeSet[0].Softirq + cpuTimeSet[0].Steal
	resp.CPUTotal = resp.CPUFree + resp.CPUUsed
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


package logic

import (
	"context"
	"errors"
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
	cpuCount, err := cpu.CountsWithContext(l.ctx, false)
	if err != nil {
		return nil, err
	}
	cpuPercent, err := cpu.PercentWithContext(l.ctx, time.Second, false)
	if err != nil {
		return nil, err
	}
	if len(cpuPercent) != 1 {
		return nil, errors.New("get cpu percent failed")
	}
	resp.CPUUsedPercent = cpuPercent[0]
	resp.CPUTotal = float64(cpuCount) * 1e9
	resp.CPUUsed = resp.CPUUsedPercent / 100 * resp.CPUTotal
	resp.CPUFree = resp.CPUTotal - resp.CPUUsed
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

package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/samber/lo"
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
	d := l.svcCtx.DockerClient
	resp = new(types.NodeTopResponse)
	// Images
	imagesSumary, err := d.ImageList(l.ctx, dockerTypes.ImageListOptions{All: true})
	if err != nil {
		return nil, err
	}
	resp.Images = lo.FlatMap(imagesSumary, func(img dockerTypes.ImageSummary, _ int) []string {
		return img.RepoTags
	})
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

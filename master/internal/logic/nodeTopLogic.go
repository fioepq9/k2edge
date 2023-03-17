package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

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

func (l *NodeTopLogic) NodeTop(req *types.NodeTopRequest) (resp *types.NodeTopResponse, err error) {
	// 判断结点是否存在
	node, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, req.Name)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("node %s does not exists", req.Name)
	}

	if !node.Status.Working {
		return nil, fmt.Errorf("node %s is not active", req.Name)
	}

	// 获取结点 WorkerURL 或 MasterURL
	resp = new(types.NodeTopResponse)
	if len(node.Roles) == 1 && lo.Contains(node.Roles, "master") {
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

	if len(node.Roles) > 0 {
		if lo.Contains(node.Roles, "worker") {
			cli := client.NewClient(node.BaseURL.WorkerURL)
			topInfo, err := cli.Node.Top(l.ctx)
			if err != nil {
				return nil, err
			}
			resp.DiskFree = topInfo.DiskFree
			resp.DiskTotal = topInfo.DiskTotal
			resp.DiskUsed = topInfo.DiskUsed
			resp.DiskUsedPercent = topInfo.DiskUsedPercent
			resp.MemoryAvailable = topInfo.MemoryAvailable
			resp.MemoryTotal = topInfo.MemoryTotal
			resp.MemoryUsed = topInfo.MemoryUsed
			resp.MemoryUsedPercent = topInfo.MemoryUsedPercent
			
		}
	}

	return resp, nil
}

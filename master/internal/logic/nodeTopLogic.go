package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"
	masterCli "k2edge/master/client"

	"github.com/samber/lo"
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

	// 结点角色只有master
	resp = new(types.NodeTopResponse)
	if len(node.Roles) == 1 && lo.Contains(node.Roles, "master") {
		cli := masterCli.NewClient(node.BaseURL.MasterURL)
		topInfo, err := cli.Node.HostTop(l.ctx)
		if err != nil {
			return nil, err
		}
		
		resp.Images = topInfo.Images
		resp.CPUFree = topInfo.CPUFree
		resp.CPUTotal = topInfo.CPUTotal
		resp.CPUUsed = topInfo.CPUUsed
		resp.CPUUsedPercent = topInfo.CPUUsedPercent
		resp.DiskFree = topInfo.DiskFree
		resp.DiskTotal = topInfo.DiskTotal
		resp.DiskUsed = topInfo.DiskUsed
		resp.DiskUsedPercent = topInfo.DiskUsedPercent
		resp.MemoryAvailable = topInfo.MemoryAvailable
		resp.MemoryTotal = topInfo.MemoryTotal
		resp.MemoryUsed = topInfo.MemoryUsed
		resp.MemoryUsedPercent = topInfo.MemoryUsedPercent

		return resp, nil
	}

	// 结点角色只为 worker， 或为 worker 和 master
	if len(node.Roles) != 0{
		cli := client.NewClient(node.BaseURL.WorkerURL)
		topInfo, err := cli.Node.Top(l.ctx)
		if err != nil {
			return nil, err
		}
		resp.CPUFree = topInfo.CPUFree
		resp.CPUTotal = topInfo.CPUTotal
		resp.CPUUsed = topInfo.CPUUsed
		resp.CPUUsedPercent = topInfo.CPUUsedPercent
		resp.Images = topInfo.Images
		resp.DiskFree = topInfo.DiskFree
		resp.DiskTotal = topInfo.DiskTotal
		resp.DiskUsed = topInfo.DiskUsed
		resp.DiskUsedPercent = topInfo.DiskUsedPercent
		resp.MemoryAvailable = topInfo.MemoryAvailable
		resp.MemoryTotal = topInfo.MemoryTotal
		resp.MemoryUsed = topInfo.MemoryUsed
		resp.MemoryUsedPercent = topInfo.MemoryUsedPercent
			
	}

	return resp, nil
}

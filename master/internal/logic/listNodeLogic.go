package logic

import (
	"context"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNodeLogic {
	return &ListNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNodeLogic) ListNode(req *types.NodeListRequest) (resp *types.NodeListResponse, err error) {
	// 获取node信息
	nodes, err := etcdutil.GetOne[types.Node](l.svcCtx.Etcd, l.ctx, "/node/" + etcdutil.SystemNamespace)

	if err != nil {
		return nil, err
	}

	resp = new(types.NodeListResponse)
	for _, node := range *nodes {
		if !req.All && !node.Status.Working {
			continue
		}
		
		n := types.NodeList{
			Name: node.Metadata.Name,
			RegisterTime: node.RegisterTime,
			URL: node.BaseURL,
		}

		if !node.Status.Working {
			n.Status = "closed"
		} else if node.Spec.Unschedulable {
			n.Status = "unschedulable"
		} else {
			n.Status = "active"
		}

		
		if len(node.Roles) == 1 {
			n.Roles = node.Roles[0]
		} else if len(node.Roles) == 2 {
			n.Roles = node.Roles[0] + "/" + node.Roles[1]
		}
		resp.NodeList = append(resp.NodeList, n)
	}

	return resp, nil
}

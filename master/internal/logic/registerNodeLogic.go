package logic

import (
	"context"
	"fmt"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterNodeLogic {
	return &RegisterNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterNodeLogic) RegisterNode(req *types.RegisterRequest) error {
	key := etcdutil.GenerateKey("node", etcdutil.SystemNamespace, req.Name)
	_, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, req.Name)

	if err != nil {
		return err
	}

	if found {
		return fmt.Errorf("node '%s' already exists", req.Name)
	}

	if len(req.Roles) == 0 {
		return fmt.Errorf("roles have not been set")
	} else if len(req.Roles) > 2 {
		return fmt.Errorf("the format of roles is wrong")
	}

	if len(req.Roles) == 2 && req.Roles[0] == req.Roles[1] {
		return fmt.Errorf("the format of roles is wrong")
	}

	for _, r := range req.Roles{
		if r != "master" && r != "worker" {
			return fmt.Errorf("the format of roles '%s' is wrong", r)
		}
	}

	if (lo.Contains(req.Roles, "master") && req.BaseURL.MasterURL == "") {
		return fmt.Errorf("master's url have not been set")
	}

	if (lo.Contains(req.Roles, "worker") && req.BaseURL.WorkerURL == "") {
		return fmt.Errorf("worker's url have not been set")
	}

	// 插入 node
	newNode := types.Node{
		Metadata: types.Metadata{
			Namespace: etcdutil.SystemNamespace,
			Kind:      "node",
			Name:      req.Name,
		},
		Roles:        req.Roles,
		BaseURL:      req.BaseURL,
		Spec: types.Spec{
			Unschedulable: false,	
		},
		RegisterTime: time.Now().Unix(),
		Status:      types.Status{
			Working: true,
		},
	}

	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, newNode)
	if err != nil {
		return err
	}
	return nil
}

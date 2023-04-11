package logic

import (
	"context"
	"errors"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeploymentLogic {
	return &ListDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDeploymentLogic) ListDeployment(req *types.ListDeploymentRequest) (resp *types.ListDeploymentResponse, err error) {
	resp = new(types.ListDeploymentResponse)

	if req.Namespace != "" {
		found, err := etcdutil.IsExistNamespace(l.svcCtx.Etcd, l.ctx, req.Namespace)
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, fmt.Errorf("namespace %s does not exist", req.Namespace)
		}
	}

	key := "/deployment"
	if req.Namespace != "" {
		key += "/" + req.Namespace 
	}

	deployments, err := etcdutil.GetOne[types.Deployment](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		if errors.Is(err, etcdutil.ErrKeyNotExist) {
			return resp, nil
		}
		return nil, err
	}

	for _, deployment := range *deployments {
		if req.Namespace == "" || deployment.Metadata.Namespace == req.Namespace {
			resp.Info = append(resp.Info, types.DeploymentSimpleInfo{
				Namespace: deployment.Metadata.Namespace,
				Name: deployment.Metadata.Name,
				CreateTime: deployment.Config.CreateTime,
				Replicas: deployment.Config.Replicas,
				Available: deployment.Status.AvailableReplicas,
			})
		}
	}
	return resp, nil
}

package logic

import (
	"context"
	"fmt"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeploymentLogic {
	return &CreateDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDeploymentLogic) CreateDeployment(req *types.CreateDeploymentRequest) (resp *types.CreateDeploymentResponse, err error) {
	if req.Deployment.Metadata.Namespace == "" {
		return nil, fmt.Errorf("deployment's namespace cannot be empty")
	}

	// 判断 deployment 的 Namespace 是否存在
	isExist, err := etcdutil.IsExistNamespace(l.svcCtx.Etcd, l.ctx, req.Deployment.Metadata.Namespace)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, fmt.Errorf("namespace %s does not exist", req.Deployment.Metadata.Namespace)
	}

	// 判断 deployment 是否已经存在
	key := etcdutil.GenerateKey("deployment", req.Deployment.Metadata.Namespace, req.Deployment.Metadata.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	if found {
		return nil, fmt.Errorf("deployment %s already exist", req.Deployment.Metadata.Name)
	}

	if req.Deployment.Config.Replicas <= 0 {
		return nil, fmt.Errorf("the replicas of deployment must be more than 0")
	}

	req.Deployment.Metadata.Kind = "deployment"
	req.Deployment.Status = types.DeploymentStatus{}

	// 创建 container 副本
	retryTimes := 3 // 创建container副本尝试次数

	createContainerRequest := &types.CreateContainerRequest{
		Container: types.Container{
			Metadata: types.Metadata{
				Namespace: req.Deployment.Metadata.Namespace,
				Kind: "container",
				Name: req.Deployment.Metadata.Name + "-" + req.Deployment.Config.Template.Name,
			},
			ContainerConfig: types.ContainerConfig{
				Deployment: req.Deployment.Metadata.Namespace + "/" + req.Deployment.Metadata.Name,
				Image: req.Deployment.Config.Template.Image,
				NodeName: req.Deployment.Config.Template.NodeName,
				Command: req.Deployment.Config.Template.Command,
				Args: req.Deployment.Config.Template.Args,
				Expose: req.Deployment.Config.Template.Expose,
				Env: req.Deployment.Config.Template.Env,
				Limit: req.Deployment.Config.Template.Limit,
				Request: req.Deployment.Config.Template.Request,
			},
		},
	}
	
	deployment := req.Deployment
	deployment.Config.CreateTime = time.Now().Unix()
	resp = new(types.CreateDeploymentResponse)
	resp.Err = make([]string, 0)
	logic := NewCreateContainerLogic(l.ctx, l.svcCtx)
	for i := 1; i <= req.Deployment.Config.Replicas; i++ {
		createContainerRequest.Container.Metadata.Name =fmt.Sprintf("%s-%s-%d",req.Deployment.Metadata.Name, req.Deployment.Config.Template.Name, i)
		var info *types.CreateContainerResponse
		for j := 1; j <= retryTimes; j++ {
			info, err = logic.CreateContainer(createContainerRequest)
			if err == nil {
				break;
			}
		}

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("creating the '%d th' container replica failed with %d retries. will try to create again later", i, retryTimes))
		} else {
			deployment.Status.Containers = append(deployment.Status.Containers, types.ContainerInfo{
				Name: createContainerRequest.Container.Metadata.Name,
				Node: info.ContainerInfo.Node,
				ContainerID: info.ContainerInfo.ContainerID,
			})
		}
	}
	deployment.Status.AvailableReplicas = req.Deployment.Config.Replicas - len(resp.Err)
	// 插入 deployment
	return resp, etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, deployment)
}

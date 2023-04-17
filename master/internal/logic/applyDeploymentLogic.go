package logic

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyDeploymentLogic {
	return &ApplyDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyDeploymentLogic) ApplyDeployment(req *types.ApplyDeploymentRequest) (resp *types.ApplyDeploymentResponse , err error) {
	resp = new(types.ApplyDeploymentResponse)
	resp.Err = make([]string, 0)
	// 判断 deployment 是否已经存在
	key := etcdutil.GenerateKey("deployment", req.Namespace, req.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("deployment %s does not exist", req.Name)
	}

	if req.Config.Replicas <= 0 {
		return nil, fmt.Errorf("the replicas of deployment must be more than 0")
	}

	deployment := types.Deployment{}
	deployment.Metadata.Namespace = req.Namespace
	deployment.Metadata.Kind = "deployment"
	deployment.Metadata.Name = req.Name
	deployment.Status = types.DeploymentStatus{}
	deployment.Config = req.Config
	deployment.Config.CreateTime = time.Now().Unix()

	// 创建 container 副本
	retryTimes := 1 // 创建container副本尝试次数

	createContainerRequest := &types.CreateContainerRequest{
		Container: types.Container{
			Metadata: types.Metadata{
				Namespace: deployment.Metadata.Namespace,
				Kind: "container",
				Name: deployment.Metadata.Name + "-" + deployment.Config.Template.Name,
			},
			ContainerConfig: types.ContainerConfig{
				Deployment: deployment.Metadata.Name,
				Image: deployment.Config.Template.Image,
				NodeName: deployment.Config.Template.NodeName,
				Command: deployment.Config.Template.Command,
				Args: deployment.Config.Template.Args,
				Expose: deployment.Config.Template.Expose,
				Env: deployment.Config.Template.Env,
				Limit: deployment.Config.Template.Limit,
				Request: deployment.Config.Template.Request,
			},
		},
	}
	

	logicC := NewCreateContainerLogic(l.ctx, l.svcCtx)
	for i := 1; i <= req.Config.Replicas; i++ {
		if req.Config.Template.Name == "" {
			createContainerRequest.Container.Metadata.Name =fmt.Sprintf("%s-%s-%d",req.Name, req.Config.Template.Image, i)
		} else {
			createContainerRequest.Container.Metadata.Name =fmt.Sprintf("%s-%s-%d",req.Name, req.Config.Template.Name, i)
		}
		var info *types.CreateContainerResponse
		for j := 1; j <= retryTimes; j++ {
			info, err = logicC.CreateContainer(createContainerRequest)
			if err == nil {
				break;
			} else {
				if strings.Contains(err.Error(), "repository does not exist") {
					return nil, err
				}
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

	deployment.Status.AvailableReplicas = deployment.Config.Replicas - len(resp.Err)

	//删除原本的 container副本
	logicG := NewGetDeploymentLogic(l.ctx, l.svcCtx)
	gresp, err := logicG.GetDeployment(&types.GetDeploymentRequest{
		Namespace: req.Namespace,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	containers := gresp.Deployment.Status.Containers
	// 删除所有容器
	for _, c := range containers {
		worker, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, c.Node)
		if err != nil {
			resp.Err = append(resp.Err, err.Error())
			continue
		}

		if !found {
			resp.Err = append(resp.Err, fmt.Sprintf("cannot find container '%s' info", c.Name))
			continue
		}

		if !worker.Status.Working {
			resp.Err = append(resp.Err, fmt.Sprintf("the node where the container '%s' is located is not active", c.Name))
			continue
		}

		// 向特定的 worker 结点发送获取conatiner信息的请求
		cli := client.NewClient(worker.BaseURL.WorkerURL)
		err = cli.Container.Stop(l.ctx, client.StopContainerRequest{
			ID:      c.ContainerID,
		})

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s', err: %s", c.Name, err))
			continue
		}

		err = cli.Container.Remove(l.ctx, client.RemoveContainerRequest{
			ID:            c.ContainerID,
		})

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s', err: %s", c.Name, err))
			continue
		}

		container1, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("container", req.Namespace, c.Name))
		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s', err: %s", c.Name, err))
			continue
		}
		c1 := (*container1)[0]

		err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("container",req.Namespace, c.Name))

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s', err: %s", c.Name, err))
			continue
		}

		err = etcdutil.NodeDeleteRequest(l.svcCtx.Etcd, l.ctx, worker.Metadata.Name, c1.ContainerConfig.Request.CPU, c1.ContainerConfig.Request.Memory)
		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s'", c.Name))
			continue
		}
	}

	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, deployment)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

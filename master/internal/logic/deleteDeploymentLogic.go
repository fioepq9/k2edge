package logic

import (
	"context"
	"fmt"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeploymentLogic {
	return &DeleteDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDeploymentLogic) DeleteDeployment(req *types.DeleteDeploymentRequest) (resp *types.DeleteDeploymentResponse, err error) {
	key := etcdutil.GenerateKey("deployment", req.Namespace, req.Name)

	// 判断 container 是否存在, 存在则获取 container 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("deployment %s does not exist", req.Name)
	}
	
 	logicG := NewGetDeploymentLogic(l.ctx, l.svcCtx)
	gresp, err := logicG.GetDeployment(&types.GetDeploymentRequest{
		Namespace: req.Namespace,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.DeleteDeploymentResponse)
	resp.Err = make([]string, 0) 
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
		container1, _ := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("container", req.Namespace, c.Name))
		c1 := (*container1)[0]
		c1.ContainerStatus.Status = "exit(0)"
		etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("container", req.Namespace, c.Name), c1)

		err = cli.Container.Stop(l.ctx, client.StopContainerRequest{
			ID:      c.ContainerID,
			Timeout: req.Timeout * int(time.Second),
		})

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while stopping the info of container '%s', err: %s", c.Name, err))
		}

		err = cli.Container.Remove(l.ctx, client.RemoveContainerRequest{
			ID:            c.ContainerID,
			RemoveVolumes: req.RemoveVolumnes,
			RemoveLinks:   req.RemoveLinks,
			Force:         req.Force,
		})

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s', err: %s", c.Name, err))
			continue
		}

		if err != nil {
			resp.Err = append(resp.Err, fmt.Sprintf("an error occurred while deleting the info of container '%s', err: %s", c.Name, err))
			continue
		}
		

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

	// 删除 etcd 中的deployment信息
	err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

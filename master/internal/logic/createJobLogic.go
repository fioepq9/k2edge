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

type CreateJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateJobLogic {
	return &CreateJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateJobLogic) CreateJob(req *types.CreateJobRequest) error {
	if req.Job.Metadata.Namespace == "" {
		return fmt.Errorf("job's namespace cannot be empty")
	}
	
	// 判断 job 的 Namespace 是否存在
	isExist, err := etcdutil.IsExistNamespace(l.svcCtx.Etcd, l.ctx, req.Job.Metadata.Namespace)
	if err != nil {
		return err
	}

	if !isExist {
		return fmt.Errorf("namespace %s does not exist", req.Job.Metadata.Namespace)
	}

	// 判断 job 是否已经存在
	key := etcdutil.GenerateKey("job", req.Job.Metadata.Namespace, req.Job.Metadata.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	if found {
		return fmt.Errorf("job %s already exist", req.Job.Metadata.Name)
	}

	if req.Job.Config.Completions <= 0 {
		return fmt.Errorf("completions must be more than 0")
	}

	req.Job.Metadata.Kind = "job"
	req.Job.Config.CreateTime = time.Now().Unix()
	job := req.Job

	err = etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, job)
	if err != nil {
		return err
	}
	if req.Job.Config.Schedule != "" {
		return nil
	}

	createContainerRequest := &types.CreateContainerRequest{
		Container: types.Container{
			Metadata: types.Metadata{
				Namespace: req.Job.Metadata.Namespace,
				Kind: "container",
				Name: req.Job.Metadata.Name + "-" + req.Job.Config.Template.Name,
			},
			ContainerConfig: types.ContainerConfig{
				Job: job.Metadata.Namespace + "/" + req.Job.Metadata.Name,
				Image: job.Config.Template.Image,
				NodeName: job.Config.Template.NodeName,
				Command: job.Config.Template.Command,
				Args: job.Config.Template.Args,
				Expose: job.Config.Template.Expose,
				Env: job.Config.Template.Env,
				Limit: job.Config.Template.Limit,
				Request: job.Config.Template.Request,
			},
		},
	}

	infos := make([]types.CreateContainerResponse, 0)
	logic := NewCreateContainerLogic(l.ctx, l.svcCtx)

	for i := 1; i <= req.Job.Config.Completions; i++ {
		if req.Job.Config.Template.Name != "" {
			createContainerRequest.Container.Metadata.Name = fmt.Sprintf("%s-%s-%d",req.Job.Metadata.Name, req.Job.Config.Template.Name, i)
		} else {
			createContainerRequest.Container.Metadata.Name = fmt.Sprintf("%s-%s-%d",req.Job.Metadata.Name, req.Job.Config.Template.Image, i)
		}
		info, errl := logic.CreateContainer(createContainerRequest)
		if info != nil {
			infos = append(infos, *info)
		}
		if errl != nil {
			for _, i := range infos {
				worker, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, i.ContainerInfo.Node)
				if err != nil {
					continue
				}

				if !found {
					continue
				}

				if !worker.Status.Working {
					continue
				}

				// 向特定的 worker 结点发送获取conatiner信息的请求
				cli := client.NewClient(worker.BaseURL.WorkerURL)
				err = cli.Container.Stop(l.ctx, client.StopContainerRequest{
					ID:      i.ContainerInfo.ContainerID,
				})

				if err != nil {
					continue
				}

				err = cli.Container.Remove(l.ctx, client.RemoveContainerRequest{
					ID:            i.ContainerInfo.ContainerID,
				})

				if err != nil {
					continue
				}

				
				container1, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("container", req.Job.Metadata.Namespace, info.ContainerInfo.Name))
				if err != nil {
					continue
				}
				c1 := (*container1)[0]

				err = etcdutil.NodeDeleteRequest(l.svcCtx.Etcd, l.ctx, worker.Metadata.Name, c1.ContainerConfig.Request.CPU, c1.ContainerConfig.Request.Memory)
				if err != nil {
					continue
				}
				err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, etcdutil.GenerateKey("container", job.Metadata.Namespace, i.ContainerInfo.Name))
				if err != nil {
					continue
				}
			}
			etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key)
			return errl
		}
	}

	return nil
}

package monitor

import (
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"
	"time"
)

func cornJob(job types.Job, s *svc.ServiceContext) error {
	key := etcdutil.GenerateKey("job", job.Metadata.Namespace, job.Metadata.Name)
	if job.Config.Schedule == "" {
		return nil
	} 

	schedule := job.Config.Schedule

	hour, minute, _ := time.Now().Clock()
	if schedule == fmt.Sprintf("%d:%02d", hour, minute) {
		createContainerRequest := &types.CreateContainerRequest{
			Container: types.Container{
				Metadata: types.Metadata{
					Namespace: job.Metadata.Namespace,
					Kind: "container",
					Name: job.Metadata.Name + "-" + job.Config.Template.Name,
				},
				ContainerConfig: types.ContainerConfig{
					Job: job.Metadata.Namespace + "/" + job.Metadata.Name,
					Image: job.Config.Template.Image,
					NodeName: job.Config.Template.NodeName,
					Command: job.Config.Template.Command,
					Args: job.Config.Template.Args,
					Expose: job.Config.Template.Expose,
					Env: job.Config.Template.Env,
					Limit:job.Config.Template.Limit,
					Request: job.Config.Template.Request,
				},
			},
		}
	
		infos := make([]types.CreateContainerResponse, 0)
		clogic := logic.NewCreateContainerLogic(s.Etcd.Ctx(), s)
	
		for i := 1; i <= job.Config.Completions; i++ {
			if job.Config.Template.Name != "" {
				createContainerRequest.Container.Metadata.Name = fmt.Sprintf("%s-%s-%d",job.Metadata.Name, job.Config.Template.Name, i)
			} else {
				createContainerRequest.Container.Metadata.Name = fmt.Sprintf("%s-%s-%d",job.Metadata.Name, job.Config.Template.Image, i)
			}
			info, errl := clogic.CreateContainer(createContainerRequest)
			if info != nil {
				infos = append(infos, *info)
			}
			if errl != nil {
				for _, i := range infos {
					worker, found, err := etcdutil.IsExistNode(s.Etcd, s.Etcd.Ctx(), i.ContainerInfo.Node)
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
					err = cli.Container.Stop(s.Etcd.Ctx(), client.StopContainerRequest{
						ID:      i.ContainerInfo.ContainerID,
					})
	
					if err != nil {
						continue
					}
	
					err = cli.Container.Remove(s.Etcd.Ctx(), client.RemoveContainerRequest{
						ID:            i.ContainerInfo.ContainerID,
					})
	
					if err != nil {
						continue
					}
	
					
					container1, err := etcdutil.GetOne[types.Container](s.Etcd, s.Etcd.Ctx(), etcdutil.GenerateKey("container", job.Metadata.Namespace, info.ContainerInfo.Name))
					if err != nil {
						continue
					}
					c1 := (*container1)[0]
	
					err = etcdutil.NodeDeleteRequest(s.Etcd, s.Etcd.Ctx(), worker.Metadata.Name, c1.ContainerConfig.Request.CPU, c1.ContainerConfig.Request.Memory)
					if err != nil {
						continue
					}
					err = etcdutil.DeleteOne(s.Etcd, s.Etcd.Ctx(), etcdutil.GenerateKey("container", job.Metadata.Namespace, i.ContainerInfo.Name))
					if err != nil {
						continue
					}
				}
				etcdutil.DeleteOne(s.Etcd, s.Etcd.Ctx(), key)
				return errl
			}
		}
	} else {
		return fmt.Errorf("not yet")
	}
	return nil
}
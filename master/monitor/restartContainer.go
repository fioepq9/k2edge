package monitor

import (
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"
)
func restart(s *svc.ServiceContext, req types.Container)(info *types.ContainerInfo, err error) {
	// 判断容器是否已经存在
	key := etcdutil.GenerateKey("container",req.Metadata.Namespace, req.Metadata.Name)
	found, err := etcdutil.IsExistKey(s.Etcd, s.Etcd.Ctx(), key)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("container %s does not exist", req.Metadata.Name)
	}
	
	// 获取目的容器在etcd中的信息以便备份恢复
	getFunc := logic.NewGetContainerLogic(s.Etcd.Ctx(), s)
	config, err := getFunc.GetContainer(&types.GetContainerRequest{
		Namespace: req.Metadata.Namespace,
		Name: req.Metadata.Name,
	})

	// 备份失败，直接终止
	if err != nil {
		return nil, fmt.Errorf("backup container %s's configuration failed", req.Metadata.Name)
	}

	// 创建容器 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.
	var worker *types.Node

	if req.Metadata.Namespace == "" {
		return nil, fmt.Errorf("container's namespace cannot be empty")
	}
	if req.Metadata.Name == "" {
		return nil, fmt.Errorf("container's name cannot be empty")
	}

	// 从 etcd 中获取需要创建容器的 worker 结点，根据在线调度算法自动获取
	worker, err = s.Worker(&req)
	if err != nil {
		return nil, err
	}

	cli := client.NewClient(worker.BaseURL.WorkerURL)
	var c types.Container
	c.Metadata = req.Metadata
	c.Metadata.Kind = "container"
	c.ContainerConfig = req.ContainerConfig
	c.ContainerStatus.Node = worker.Metadata.Name

	expose := make([]client.ExposedPort, 0)
	for _, e := range c.ContainerConfig.Expose {
		expose = append(expose, client.ExposedPort{
			Port:     e.Port,
			Protocol: e.Protocol,
			HostPort: e.HostPort,
		})
	}

	// 访问 worker 结点并创建容器
	res, err := cli.Container.Create(s.Etcd.Ctx(), client.CreateContainerRequest{
		ContainerName: c.Metadata.Name,
		Config: client.ContainerConfig{
			Deployment: req.ContainerConfig.Deployment,
			Job:	req.ContainerConfig.Job,
			Image:    c.ContainerConfig.Image,
			NodeName: c.ContainerConfig.NodeName,
			Command:  c.ContainerConfig.Command,
			Args:     c.ContainerConfig.Args,
			Expose:   expose,
			Env:      c.ContainerConfig.Env,
			Limit:    client.ContainerLimit{
				CPU: c.ContainerConfig.Limit.CPU,
				Memory: c.ContainerConfig.Limit.Memory,
			},
			Request:    client.ContainerRequest{
				CPU: c.ContainerConfig.Request.CPU,
				Memory: c.ContainerConfig.Request.Memory,
			},
		},
	})

	if err != nil {
		return nil, err
	}
	c.ContainerStatus.ContainerID = res.ID
	c.ContainerStatus.Status = "running"
	// 在worker节点创建容器成功 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// 将容器信息写入etcd
	err = etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, c)
	if err != nil { // 写入etcd失败
		// 删除worker中新创建的容器
		errd := cli.Container.Stop(s.Etcd.Ctx(), client.StopContainerRequest{
			ID:      c.ContainerStatus.ContainerID,
		})
	
		if errd != nil {
			return nil, fmt.Errorf("apply container failed, recover failed : %s", err.Error())
		}
	
		errd = cli.Container.Remove(s.Etcd.Ctx(), client.RemoveContainerRequest{
			ID:            c.ContainerStatus.ContainerID,
		})
	
		if errd != nil {
			return nil, fmt.Errorf("apply container failed, recover failed : %s", err.Error())
		}
		
		return nil, err
	}

	// 写入etcd成功后，删除原本的container
	// 删除worker中新创建的容器
	err = cli.Container.Stop(s.Etcd.Ctx(), client.StopContainerRequest{
		ID:      config.Container.ContainerStatus.ContainerID,
	})

	if err != nil {
		return nil, fmt.Errorf("apply container failed, delete orignal container failed : %s", err.Error())
	}

	err = cli.Container.Remove(s.Etcd.Ctx(), client.RemoveContainerRequest{
		ID:            config.Container.ContainerStatus.ContainerID,
	})

	if err != nil {
		return nil, fmt.Errorf("apply container failed, delete orignal container failed : %s", err.Error())
	}
	return &types.ContainerInfo{
		Name: c.Metadata.Name,
		Node: c.ContainerStatus.Node,
		ContainerID: c.ContainerStatus.ContainerID,
	}, nil
}

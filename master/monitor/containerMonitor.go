package monitor

import (
	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
)

func contianerEvent(event types.EventInfo, s *svc.ServiceContext) error {
	key := etcdutil.GenerateKey("container", event.Container.Namespace, event.Container.Name)
	containers, err := etcdutil.GetOne[types.Container](s.Etcd, s.Etcd.Ctx(), key)
	if err != nil {
		return err
	}
	container := (*containers)[0]

	if container.ContainerStatus.Status == "processing" {
		return nil
	}

	backup := container
	container.ContainerStatus.Status = "processing"
	etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, container)

	
	if event.Action == "start" && container.ContainerStatus.Status != "running" {
		container.ContainerStatus.Status = "running"
		err = etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, container)
		if err != nil {
			return  err
		}
		return nil
	}

	if event.ExitCode == "0" {
		if container.ContainerStatus.Status != "exit(0)" {
			container.ContainerStatus.Status = "exit(0)"
			err = etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, container)
			if err != nil {
				return  err
			}
		}
		return nil
	}

	
	_, err = restart(s, container)
	if err != nil {
		etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, backup)
		return  err
	}
	return nil
}

//异常状态处理
func containerStatus(container types.Container, s *svc.ServiceContext) error {
	if container.ContainerStatus.Status == "processing" {
		return nil
	}

	key := etcdutil.GenerateKey("container", container.Metadata.Namespace, container.Metadata.Name)
	backup := container
	container.ContainerStatus.Status = "processing"
	etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, container)
	_, err := restart(s, container)
	if err != nil {
		etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), key, backup)
		return  err
	}
	return nil
}
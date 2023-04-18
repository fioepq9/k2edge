package monitor

import (
	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"strings"

	"github.com/samber/lo"
)

func deploymentEvent(event types.EventInfo, s *svc.ServiceContext) error {
	ckey := etcdutil.GenerateKey("container", event.Container.Namespace, event.Container.Name)
	containers, err := etcdutil.GetOne[types.Container](s.Etcd, s.Etcd.Ctx(), ckey)
	if err != nil {
		return err
	}

	container := (*containers)[0]
	if event.ExitCode == "0" {
		container.ContainerStatus.Status = "exit(0)"
		err = etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), ckey, container)
		if err != nil {
			return  err
		}
		return nil
	}

	info, err := restart(s, container)
	if err != nil {
		return err
	}
	t := strings.Split(container.ContainerConfig.Deployment, "/")
	dnamespace := t[0]
	dname := t[1]
	dkey := etcdutil.GenerateKey("deployment", dnamespace, dname)
	deployments, err := etcdutil.GetOne[types.Deployment](s.Etcd, s.Etcd.Ctx(), dkey)
	if err != nil {
		return err
	}
	deployment := (*deployments)[0]
	deployment.Status.Containers = lo.Filter(deployment.Status.Containers, func(item types.ContainerInfo, index int) bool {
		return item.ContainerID != info.ContainerID
	})
	deployment.Status.Containers = append(deployment.Status.Containers, *info)

	return etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), dkey, deployment)
}
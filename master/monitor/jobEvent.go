package monitor

import (
	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"strings"

	"k2edge/master/internal/logic"
)

func jobEvent(event types.EventInfo, s *svc.ServiceContext) error {
	ckey := etcdutil.GenerateKey("container", event.Container.Namespace, event.Container.Name)
	containers, err := etcdutil.GetOne[types.Container](s.Etcd, s.Etcd.Ctx(), ckey)
	if err != nil {
		return err
	}

	container := (*containers)[0]

	t := strings.Split(container.ContainerConfig.Job, "/")
	jnamespace := t[0]
	jname := t[1]
	jkey := etcdutil.GenerateKey("job", jnamespace, jname)
	jobs, err := etcdutil.GetOne[types.Job](s.Etcd, s.Etcd.Ctx(), jkey)
	if err != nil {
		return err
	}
	job := (*jobs)[0]

	if event.ExitCode == "0" {
		job.Succeeded++
		err = etcdutil.PutOne(s.Etcd, s.Etcd.Ctx(), jkey, job)
		if err != nil {
			return err
		}
		dlogic := logic.NewDeleteContainerLogic(s.Etcd.Ctx(),s)
		err = dlogic.DeleteContainer(&types.DeleteContainerRequest{
			Namespace: container.Metadata.Namespace,
			Name: container.Metadata.Name,
		})
		if err != nil {
			return err
		}
		return nil
	}

	_, err = restart(s, container)
	if err != nil {
		return  err
	}
	return nil
}
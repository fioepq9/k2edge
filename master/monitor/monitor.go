package monitor

import (
	"context"
	"errors"
	"fmt"
	"k2edge/etcdutil"
	masterCli "k2edge/master/client"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"
	"time"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

func EventMonitor(svcCtx *svc.ServiceContext) {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		event, ekey, err := etcdutil.GetOneKV[types.EventInfo](svcCtx.Etcd, svcCtx.Etcd.Ctx(), "/events")
		if err != nil {
			if !errors.Is(err, etcdutil.ErrKeyNotExist) {
				logx.Error(err)
			}
			continue
		}
		err = etcdutil.DeleteOne(svcCtx.Etcd, svcCtx.Etcd.Ctx(), ekey)
		if err != nil {
			logx.Error(err)
			continue
		}
		if event.Action != "start" && event.Container.Deployment != "" {
			err = deploymentEvent(*event, svcCtx)
		} else if event.Action != "start" && event.Container.Job != "" {
			err = jobEvent(*event, svcCtx)
		} else {
			err = contianerEvent(*event, svcCtx)
		}
		if err == nil {
			continue
		}
		logx.Error(err)

		event.Times++
		if event.Times < 3 {
			err = etcdutil.PutOne(svcCtx.Etcd, svcCtx.Etcd.Ctx(), fmt.Sprintf("/events/%d-%s", time.Now().Unix(), event.Container.Id), *event)
		} else {
			ckey := etcdutil.GenerateKey("container", event.Container.Namespace, event.Container.Name)
			containers, err := etcdutil.GetOne[types.Container](svcCtx.Etcd, svcCtx.Etcd.Ctx(), ckey)
			if err != nil {
				logx.Error(err)
				continue
			}
			container := (*containers)[0]
			if event.Action == "die" {
				container.ContainerStatus.Status = fmt.Sprintf("exit(%s)", event.ExitCode)
			} else if event.Action == "start" {
				container.ContainerStatus.Status = "running"
			}

			errp := etcdutil.PutOne(svcCtx.Etcd, svcCtx.Etcd.Ctx(), ckey, container)
			logx.Error(errp)
		}
		if err != nil {
			logx.Error(err)
		}
	}
}

func StatusMonitor(svcCtx *svc.ServiceContext) {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		containers, err := etcdutil.GetOne[types.Container](svcCtx.Etcd, svcCtx.Etcd.Ctx(), "/container")
		if err != nil {
			if !errors.Is(err, etcdutil.ErrKeyNotExist) {
				logx.Error(err)
			}
			continue
		}

		for _, container := range *containers {
			if (container.ContainerStatus.Status != "exit(0)") &&
				container.ContainerStatus.Status != "running" {
				if container.ContainerConfig.Deployment != "" {
					err = deploymentStatus(container, svcCtx)
				} else if container.ContainerConfig.Job != "" {
					err = jobStatus(container, svcCtx)
				} else {
					err = containerStatus(container, svcCtx)
				}
				if err != nil {
					logx.Error(err)
				}
			}
		}
	}
}

func NodeMonitor(svcCtx *svc.ServiceContext) {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		nodes, err := etcdutil.GetOne[types.Node](svcCtx.Etcd, svcCtx.Etcd.Ctx(), "/node/"+etcdutil.SystemNamespace)
		if err != nil {
			if !errors.Is(err, etcdutil.ErrKeyNotExist) {
				logx.Error(err)
			}
			continue
		}

		for _, node := range *nodes {
			ctx, cancel := context.WithTimeout(svcCtx.Etcd.Ctx(), 5*time.Second)
			//fmt.Printf("%#v\n", node)
			if lo.Contains(node.Roles, "worker") {
				cli := client.NewClient(node.BaseURL.WorkerURL)
				resp, err := cli.Node.Version(ctx)
				//fmt.Printf("%s %#v\n", err, "worker")
				if err != nil || resp == nil {
					nodeStatus(node, svcCtx, false)
				} else {
					nodeStatus(node, svcCtx, true)
				}
			}

			if lo.Contains(node.Roles, "master") {
				cli := masterCli.NewClient(node.BaseURL.MasterURL)
				resp, err := cli.Cluster.Info(ctx)
				//fmt.Printf("%s %#v\n", err, "master")
				if err != nil || resp == nil {
					nodeStatus(node, svcCtx, false)
				} else {
					nodeStatus(node, svcCtx, true)
				}
			}
			cancel()
		}

	}
}

func CornjobMonitor(svcCtx *svc.ServiceContext) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		jobs, err := etcdutil.GetOne[types.Job](svcCtx.Etcd, svcCtx.Etcd.Ctx(), "/job")
		if err != nil {
			if !errors.Is(err, etcdutil.ErrKeyNotExist) {
				logx.Error(err)
			}
			continue
		}

		for _, job := range *jobs {
			if job.Config.Schedule != "" {
				jkey := etcdutil.GenerateKey("job", job.Metadata.Namespace, job.Metadata.Name)
				err := etcdutil.DeleteOne(svcCtx.Etcd, svcCtx.Etcd.Ctx(), jkey)
				if err != nil {
					logx.Error(err)
					continue
				}
				cornJob(job, svcCtx)
				etcdutil.PutOne(svcCtx.Etcd, svcCtx.Etcd.Ctx(), jkey, job)
			}
		}
	}
}

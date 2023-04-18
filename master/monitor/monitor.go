package monitor

import (
	"errors"
	"fmt"
	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func Monite(svcCtx *svc.ServiceContext) {
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
		if event.Container.Deployment != "" {
			err = deploymentEvent(*event, svcCtx)
		} else if event.Container.Job != "" {
			err = jobEvent(*event, svcCtx)
		} else {
			err = contianerEvent(*event, svcCtx)
		}
		if err != nil {
			logx.Error(err)
		}

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
			container.ContainerStatus.Status = fmt.Sprintf("exit(%s)", event.ExitCode)
			errp := etcdutil.PutOne(svcCtx.Etcd, svcCtx.Etcd.Ctx(), ckey, container)
			logx.Error(errp)
		}
		if err != nil {
			logx.Error(err)
		}
	}
}

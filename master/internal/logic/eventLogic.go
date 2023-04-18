package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventLogic {
	return &EventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventLogic) Event(req *types.EventRequest) error {
	fmt.Printf("%#v\n",req)
	event := types.EventInfo{}
	//判断事件是否又必要存入到etcd中
	msg := req.Message
	containerID := msg.ID
	//判断是否是异常退出
	attributes := make(map[string]string)
	for _, a := range msg.Actor.Attributes {
		s := strings.Split(a, "::")
		attributes[s[0]] = s[1]
	}
	event.ExitCode = attributes["exitCode"]

	// 获取所有container
	keyc := "/container"
	containers, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, keyc)
	if err != nil {
		if errors.Is(err, etcdutil.ErrKeyNotExist) {
			return nil
		}
		return err
	}

	needWrite := false
	for _, container := range *containers {
		if container.ContainerStatus.ContainerID == containerID {
			event.Container.Namespace = container.Metadata.Namespace
			event.Container.Name = container.Metadata.Name
			event.Container.Id = container.ContainerStatus.ContainerID
			event.Container.Node = container.ContainerStatus.Node
			event.Container.Deployment = container.ContainerConfig.Deployment
			event.Container.Job = container.ContainerConfig.Job
			needWrite = true
			break
		}
	}

	//不需要写入etcd
	if !needWrite {
		return nil
	}

	// event信息赋值
	event.Times = 0
	event.Action = msg.Action
	event.Time = msg.Time
	if event.Time == 0 {
		event.Time = time.Now().Unix()
	}

	// 写入etcd
	key := fmt.Sprintf("/events/%d-%s", event.Time, event.Container.Id)

	return etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, key, event)
}

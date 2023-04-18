package event

import (
	"context"
	"errors"
	"k2edge/etcdutil"
	"k2edge/worker/internal/svc"
	"time"

	"k2edge/worker/client"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EventSubcriber struct {
	svcCtx *svc.ServiceContext
}

func NewEventSubcriber(svcCtx *svc.ServiceContext) *EventSubcriber {
	return &EventSubcriber{
		svcCtx: svcCtx,
	}
}

func (s *EventSubcriber) Subcribe() error {
	d := s.svcCtx.Docker
	msgReceiver, errReceiver := d.Events(context.Background(), types.EventsOptions{})
	for {
		select {
		case msg := <-msgReceiver:
			err := s.SendMessage(msg)
			if err != nil {
				logx.Errorf("EventSubcriber SendMessage failed, error=%s", err)
			}
		case err := <-errReceiver:
			return err
		}
	}
}

func (s *EventSubcriber) SendMessage(msg events.Message) error {
	if msg.Action != "die" {
		return nil
	}

	server, err := getServer(s.svcCtx.Config.Etcd.Endpoints[0])
	if err != nil {
		return err
	}

	cli := client.NewClient(server)
	attribute := []string{}
	for key, value := range msg.Actor.Attributes {
		attribute = append(attribute, key + "::" + value)
	}

	err = cli.Container.Event(s.svcCtx.Etcd.Ctx(), client.EventRequest{
		Message: client.Message{
			Status: msg.Status,
			ID:		msg.ID,
			From:	msg.From,
			Type:   msg.Type,
			Action: msg.Action,
			Actor:	client.Actor{
				ID: msg.Actor.ID,
				Attributes: attribute,
			},
			Scope:  msg.Scope,    
			Time:	msg.Time,
			TimeNano: msg.TimeNano,
		},
	})
	if err != nil {
		return err
	}

	//fmt.Printf("%+v\n", msg)
	return nil
}

func getServer(endPoint string) (server string, err error) {
	config := clientv3.Config{
		Endpoints:   []string{endPoint},
		DialTimeout: 5 * time.Second,
	}

	etcd, err := clientv3.New(config)
	if err != nil {
		return "", err
	}
	// 获取node信息
	nodes, err := etcdutil.GetOne[client.Node](etcd, context.Background(), "/node/" + etcdutil.SystemNamespace)

	if err != nil {
		return "", err
	}

	for _, node := range *nodes {
		if node.Status.Working && lo.Contains(node.Roles, "master") {
			return node.BaseURL.MasterURL, nil
		}
	}
	return "", errors.New("")
}
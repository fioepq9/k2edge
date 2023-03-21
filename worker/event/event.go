package event

import (
	"context"
	"fmt"
	"k2edge/worker/internal/svc"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/zeromicro/go-zero/core/logx"
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
	fmt.Printf("%+v\n", msg)
	return nil
}

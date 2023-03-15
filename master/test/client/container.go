package client

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/imroc/req/v3"
)

type containerAPI struct {
	opt *ClientOption
	req *req.Client
}

func (c containerAPI) Exec(ctx context.Context, req ExecContainerRequest) (io.ReadWriteCloser, error) {
	vals := make(url.Values)
	rv := reflect.ValueOf(req)
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		tagslice := strings.Split(rt.Field(i).Tag.Get("form"), ",")
		if len(tagslice) == 0 {
			continue
		}
		tagName := tagslice[0]
		if !rv.Field(i).IsZero() {
			vals.Add(tagName, fmt.Sprint(rv.Field(i).Interface()))
		}
	}
	fmt.Println()
	conn, _, err := websocket.DefaultDialer.DialContext(
		ctx,
		fmt.Sprintf("%s/container/exec?%s", c.opt.WebsocketBaseURL(), vals.Encode()),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &websocketSession{
		ws: conn,
	}, nil
}

func (c containerAPI) Attach(ctx context.Context, req AttachContainerRequest) (io.ReadWriteCloser, error) {
	vals := make(url.Values)
	rv := reflect.ValueOf(req)
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		tagslice := strings.Split(rt.Field(i).Tag.Get("form"), ",")
		if len(tagslice) == 0 {
			continue
		}
		tagName := tagslice[0]
		if !rv.Field(i).IsZero() {
			vals.Add(tagName, fmt.Sprint(rv.Field(i).Interface()))
		}
	}
	conn, _, err := websocket.DefaultDialer.DialContext(
		ctx,
		fmt.Sprintf("%s/container/attach?%s", c.opt.WebsocketBaseURL(), vals.Encode()),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &websocketSession{
		ws: conn,
	}, nil
}

type websocketSession struct {
	ws *websocket.Conn
}

func (s *websocketSession) Read(p []byte) (n int, err error) {
	_, rd, err := s.ws.NextReader()
	if err != nil {
		return 0, err
	}
	return rd.Read(p)
}

func (s *websocketSession) Write(p []byte) (n int, err error) {
	wt, err := s.ws.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return 0, err
	}
	defer wt.Close()
	return wt.Write(p)
}

func (s *websocketSession) Close() error {
	return s.ws.Close()
}

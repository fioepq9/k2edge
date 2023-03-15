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

type containers struct {
	opt *ClientOption
	cli *req.Client
}

type containerExecSession struct {
	ws *websocket.Conn
}

func (s *containerExecSession) Read(p []byte) (n int, err error) {
	_, rd, err := s.ws.NextReader()
	if err != nil {
		return 0, err
	}
	return rd.Read(p)
}

func (s *containerExecSession) Write(p []byte) (n int, err error) {
	wt, err := s.ws.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return 0, err
	}
	defer wt.Close()
	return wt.Write(p)
}

func (s *containerExecSession) Close() error {
	return s.ws.Close()
}

func (c containers) Create(ctx context.Context, req CreateContainerRequest) (resp *CreateContainerResponse, err error) {
	err = c.cli.
		Post("/container/create").
		SetBodyJsonMarshal(req).
		Do(ctx).Into(&resp)
	return
}

func (c containers) Remove(ctx context.Context, req RemoveContainerRequest) error {
	return c.cli.Post("/container/remove").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Stop(ctx context.Context, req StopContainerRequest) error {
	return c.cli.Post("/container/stop").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Start(ctx context.Context, req StartContainerRequest) error {
	return c.cli.Post("/container/start").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Status(ctx context.Context, req ContainerStatusRequest) (resp *ContainerStatusResponse, err error) {
	err = c.cli.
		Get("/container/status").
		AddQueryParam("id", req.ID).
		Do(ctx).Into(&resp)
	return
}

func (c containers) List(ctx context.Context, req ListContainersRequest) (resp *ListContainersResponse, err error) {
	resp = new(ListContainersResponse)
	params := make(map[string]interface{})
	rv := reflect.ValueOf(req)
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		tagslice := strings.Split(rt.Field(i).Tag.Get("form"), ",")
		if len(tagslice) == 0 {
			continue
		}
		tagName := tagslice[0]
		if !rv.Field(i).IsZero() {
			params[tagName] = rv.Field(i).Interface()
		}
	}
	cli := c.cli.Get("/container/list")
	for k, v := range params {
		cli = cli.AddQueryParam(k, fmt.Sprint(v))
	}
	err = cli.Do(ctx).Into(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c containers) Exec(ctx context.Context, req ExecRequest) (io.ReadWriteCloser, error) {
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
		fmt.Sprintf("ws://%s:%d/container/exec?%s", c.opt.host, c.opt.port, vals.Encode()),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &containerExecSession{
		ws: conn,
	}, nil
}

func (c containers) Attach(ctx context.Context, req AttachRequest) error {
	return c.cli.Post("/container/attach").SetBodyJsonMarshal(req).Do(ctx).Err
}

package client

import (
	"context"
	"fmt"
	"io"
	"k2edge/master/internal/types"
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

func (c containerAPI) Create(ctx context.Context, req types.CreateContainerRequest) error {
	return c.req.Post("/container/create").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containerAPI) Apply(ctx context.Context, req types.ApplyContainerRequest) error {
	return c.req.Post("/container/apply").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containerAPI) Get(ctx context.Context, req types.GetContainerRequest) (resp *types.GetContainerResponse, err error) {
	err = c.req.
		Get("/container/get").
		AddQueryParam("namespace", req.Namespace).
		AddQueryParam("name", req.Name).
		Do(ctx).Into(&resp)
	return
}

func (c containerAPI) List(ctx context.Context, req types.ListContainerRequest) (resp *types.ListContainerResponse, err error) {
	err = c.req.
		Get("/container/list").
		AddQueryParam("namespace", req.Namespace).
		Do(ctx).Into(&resp)
	return
}

func (c containerAPI) Delete(ctx context.Context, req types.DeleteContainerRequest) error {
	return c.req.Post("/container/delete").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containerAPI) Exec(ctx context.Context, req  types.ExecContainerRequest) (io.ReadWriteCloser, error) {
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

func (c containerAPI) Attach(ctx context.Context, req types.AttachContainerRequest) (io.ReadWriteCloser, error) {
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

func (c containerAPI) Logs(ctx context.Context, req types.LogsContainerRequest) (io.ReadCloser, error) {
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
		fmt.Sprintf("%s/container/logs?%s", c.opt.WebsocketBaseURL(), vals.Encode()),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &websocketReadSession{
		ws: conn,
	}, nil
}

func (c containerAPI) Migrate(ctx context.Context, req types.MigrateContainerRequest) error {
	return c.req.Post("/container/migrate").SetBodyJsonMarshal(req).Do(ctx).Err
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

type websocketReadSession struct {
	ws *websocket.Conn
}

func (s *websocketReadSession) Read(p []byte) (n int, err error) {
	_, rd, err := s.ws.NextReader()
	if err != nil {
		return 0, err
	}
	return rd.Read(p)
}

func (s *websocketReadSession) Close() error {
	return s.ws.Close()
}

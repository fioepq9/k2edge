package client

import (
	"context"
	"fmt"
	"k2edge/worker/internal/types"
	"reflect"
	"strings"

	"github.com/imroc/req/v3"
)

func (c *Client) Containers() containers {
	return containers{
		cli: c.Client,
	}
}

type containers struct {
	cli *req.Client
}

func (c containers) Run(ctx context.Context, req types.RunContainerRequest) error {
	return c.cli.Post("/container/run").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Remove(ctx context.Context, req types.RemoveContainerRequest) error {
	return c.cli.Post("/container/remove").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Stop(ctx context.Context, req types.StopContainerRequest) error {
	return c.cli.Post("/container/stop").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Start(ctx context.Context, req types.StartContainerRequest) error {
	return c.cli.Post("/container/start").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Status(ctx context.Context, req types.ContainerStatusRequest) (resp *types.ContainerStatusResponse, err error) {
	resp = new(types.ContainerStatusResponse)
	err = c.cli.Get("/container/status").AddQueryParam("id", req.ID).Do(ctx).Into(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c containers) List(ctx context.Context, req types.ListContainersRequest) (resp *types.ListContainersResponse, err error) {
	resp = new(types.ListContainersResponse)
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

func (c containers) Exec(ctx context.Context, req types.ExecRequest) error {
	return c.cli.Post("/container/exec").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (c containers) Attach(ctx context.Context, req types.AttachRequest) error {
	return c.cli.Post("/container/attach").SetBodyJsonMarshal(req).Do(ctx).Err
}
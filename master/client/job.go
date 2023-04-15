package client

import (
	"context"
	"k2edge/master/internal/types"

	"github.com/imroc/req/v3"
)

type jobAPI struct {
	opt *ClientOption
	req *req.Client
}

func (j jobAPI) Create(ctx context.Context, req types.CreateJobRequest) error {
	return j.req.Post("/job/create").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (j jobAPI) Get(ctx context.Context, req types.GetJobRequest) (resp *types.GetJobResponse, err error) {
	err = j.req.
		Get("/job/get").
		AddQueryParam("namespace", req.Namespace).
		AddQueryParam("name", req.Name).
		Do(ctx).Into(&resp)
	return
}

func (j jobAPI) List(ctx context.Context, req types.ListJobRequest) (resp *types.ListJobResponse, err error) {
	err = j.req.
		Get("/job/list").
		AddQueryParam("namespace", req.Namespace).
		Do(ctx).Into(&resp)
	return
}

func (j jobAPI) Delete(ctx context.Context, req types.DeleteJobRequest) error {
	return j.req.Post("/job/delete").SetBodyJsonMarshal(req).Do(ctx).Err
}

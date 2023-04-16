package client

import (
	"context"
	"k2edge/master/internal/types"

	"github.com/imroc/req/v3"
)

type deploymentAPI struct {
	opt *ClientOption
	req *req.Client
}

func (d deploymentAPI) Create(ctx context.Context, req types.CreateDeploymentRequest) (resp *types.CreateDeploymentResponse, err error) {
	err = d.req.Post("/deployment/create").SetBodyJsonMarshal(req).Do(ctx).Into(&resp)
	return
}

func (d deploymentAPI) Get(ctx context.Context, req types.GetDeploymentRequest) (resp *types.GetDeploymentResponse, err error) {
	err = d.req.
		Get("/deployment/get").
		AddQueryParam("namespace", req.Namespace).
		AddQueryParam("name", req.Name).
		Do(ctx).Into(&resp)
	return
}

func (d deploymentAPI) List(ctx context.Context, req types.ListDeploymentRequest) (resp *types.ListDeploymentResponse, err error) {
	err = d.req.
		Get("/deployment/list").
		AddQueryParam("namespace", req.Namespace).
		Do(ctx).Into(&resp)
	return
}

func (d deploymentAPI) Delete(ctx context.Context, req types.DeleteDeploymentRequest) error {
	return d.req.Post("/deployment/delete").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (d deploymentAPI) Apply(ctx context.Context, req types.ApplyDeploymentRequest)(resp *types.ApplyDeploymentResponse, err error) {
	err = d.req.Post("/deployment/apply").SetBodyJsonMarshal(req).Do(ctx).Into(&resp)
	return
}

func (d deploymentAPI) Scale(ctx context.Context, req types.ScaleRequest) error {
	return d.req.Post("/deployment/scale").SetBodyJsonMarshal(req).Do(ctx).Err
}

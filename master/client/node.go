package client

import (
	"context"

	"k2edge/master/internal/types"

	"github.com/imroc/req/v3"
)

type nodeAPI struct {
	opt *ClientOption
	req *req.Client
}

func (n nodeAPI) Register(ctx context.Context, req *types.RegisterRequest) error {
	return n.req.Post("/node/register").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n nodeAPI) List(ctx context.Context, req types.NodeListRequest) (resp *types.NodeListResponse, err error) {
	all := "false"
	if req.All {
		all = "true"
	}

	err = n.req.
		Get("/node/list").
		AddQueryParam("all", all).
		Do(ctx).Into(&resp)
	return
}

func (n nodeAPI) Top(ctx context.Context, req types.NodeTopRequest) (resp *types.NodeListResponse, err error) {
	err = n.req.
		Get("/node/top").
		AddQueryParam("name", req.Name).
		Do(ctx).Into(&resp)
	return
}

func (n nodeAPI) cordon(ctx context.Context, req types.CordonRequest) error {
	return n.req.Post("/node/cordon").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n nodeAPI) uncordon(ctx context.Context, req types.UncordonRequest) error {
	return n.req.Post("/node/uncordon").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n nodeAPI) drain(ctx context.Context, req types.DrainRequest) error {
	return n.req.Post("/node/drain").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n nodeAPI) delete(ctx context.Context, req types. DeleteRequest) error {
	return n.req.Post("/node/delete").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n nodeAPI) HostTop(ctx context.Context) (resp *types.NodeTopResponse, err error) {
	err = n.req.
		Get("/node/hostTop").
		Do(ctx).Into(&resp)
	return
}

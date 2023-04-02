package client

import (
	"context"

	"github.com/imroc/req/v3"
	"k2edge/master/internal/types"
)

type namespaceAPI struct {
	opt *ClientOption
	req *req.Client
}

func (n namespaceAPI) namespaceCreate(ctx context.Context, req types.CreateContainerRequest) error {
	return n.req.Post("/namespace/create").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n namespaceAPI) namespaceGet(ctx context.Context, req types.GetContainerRequest) (resp *types.GetContainerResponse, err error) {
	err = n.req.
		Get("/namespace/get").
		Do(ctx).Into(&resp)
	return
}

func (n namespaceAPI) namespaceList(ctx context.Context, req types.ListContainerRequest) (resp *types.ListContainerResponse, err error) {
	err = n.req.
		Get("/namespace/list").
		Do(ctx).Into(&resp)
	return
}

func (n namespaceAPI) namespaceDelete(ctx context.Context, req types.DeleteContainerRequest) error {
	return n.req.Post("/namespace/delete").SetBodyJsonMarshal(req).Do(ctx).Err
}
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

func (n namespaceAPI) NamespaceCreate(ctx context.Context, req types.CreateNamespaceRequest) error {
	return n.req.Post("/namespace/create").SetBodyJsonMarshal(req).Do(ctx).Err
}

func (n namespaceAPI) NamespaceGet(ctx context.Context, req types.GetNamespaceRequest) (resp *types.GetNamespaceResponse, err error) {
	err = n.req.
		Get("/namespace/get").
		AddQueryParam("name", req.Name).
		Do(ctx).Into(&resp)
	return
}

func (n namespaceAPI) NamespaceList(ctx context.Context, req types.ListNamespaceRequest) (resp *types.ListNamespaceResponse, err error) {
	all := "false"
	if req.All {
		all = "true"
	}
	
	err = n.req.
		Get("/namespace/list").
		AddQueryParam("all", all).
		Do(ctx).Into(&resp)
	return
}

func (n namespaceAPI) NamespaceDelete(ctx context.Context, req types.DeleteNamespaceRequest) error {
	return n.req.Post("/namespace/delete").SetBodyJsonMarshal(req).Do(ctx).Err
}
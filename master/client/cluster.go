package client

import (
	"context"
	"k2edge/master/internal/types"
	"github.com/imroc/req/v3"
)

type clusterAPI struct {
	opt *ClientOption
	req *req.Client
}

func (c clusterAPI) Info(ctx context.Context) (resp *types.ClusterInfoResponse, err error) {
	err = c.req.
		Get("/cluster/info").
		Do(ctx).Into(&resp)
	return
}
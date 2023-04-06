package client

import (
	"context"

	"github.com/imroc/req/v3"
	"k2edge/master/internal/types"
)

type nodeAPI struct {
	opt *ClientOption
	req *req.Client
}

func (n nodeAPI) Top(ctx context.Context) (resp *types.NodeTopResponse, err error) {
	err = n.req.
		Get("/node/hostTop").
		Do(ctx).Into(&resp)
	return
}
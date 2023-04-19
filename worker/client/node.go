package client

import (
	"context"
	"k2edge/worker/internal/types"

	"github.com/imroc/req/v3"
)

type nodeAPI struct {
	opt *ClientOption
	req *req.Client
}

func (n nodeAPI) Top(ctx context.Context) (resp *types.NodeTopResponse, err error) {
	err = n.req.
		Get("/node/top").
		Do(ctx).Into(&resp)
	return
}
func (n nodeAPI) Version(ctx context.Context) (resp *types.VersionResponse, err error) {
	err = n.req.
		Get("/version").
		Do(ctx).Into(&resp)
	return
}

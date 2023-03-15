package client

import (
	"context"
	"k2edge/worker/internal/types"

	"github.com/imroc/req/v3"
)

type nodes struct {
	opt *ClientOption
	cli *req.Client
}

func (n nodes) Top(ctx context.Context) (resp *types.NodeTopResponse, err error) {
	err = n.cli.
		Get("/node/top").
		Do(ctx).Into(&resp)
	return
}

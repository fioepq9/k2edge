package client

import (
	"context"
	"k2edge/worker/internal/types"

	"github.com/imroc/req/v3"
)

func (c *Client) Nodes() nodes {
	return nodes{
		cli: c.Client,
	}
}

type nodes struct {
	cli *req.Client
}

func (n nodes) Top(ctx context.Context) (resp *types.NodeTopResponse, err error) {
	err = n.cli.
		Get("/node/top").
		Do(ctx).Into(&resp)
	return
}

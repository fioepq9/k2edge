package client

import (
	"encoding/json"
	"fmt"
	"k2edge/worker/internal/types"
	"time"

	"github.com/imroc/req/v3"
)

type Client struct {
	opt *ClientOption
	c   *req.Client

	Container containers
	Node      nodes
}

type ClientOption struct {
	host string
	port int
}

type Option func(*ClientOption)

func WithHost(host string) Option {
	return func(co *ClientOption) {
		co.host = host
	}
}

func WithPort(port int) Option {
	return func(co *ClientOption) {
		co.port = port
	}
}

func NewClient(opt ...Option) *Client {
	var co ClientOption
	for _, o := range opt {
		o(&co)
	}
	if co.host == "" {
		co.host = "localhost"
	}
	if co.port == 0 {
		co.port = 8888
	}
	var cli Client
	cli.opt = &co
	cli.c = req.C().
		SetBaseURL(fmt.Sprintf("http://%s:%d", co.host, co.port)).
		SetCommonRetryCount(2).
		SetCommonRetryBackoffInterval(time.Second, 5*time.Second).
		AddCommonRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode >= 500
		}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = fmt.Errorf("bad response, raw content:\n%s", resp.Dump())
				return nil
			}
			return nil
		}).
		SetResponseBodyTransformer(func(rawBody []byte, req *req.Request, resp *req.Response) (transformedBody []byte, err error) {
			var r types.Response
			err = json.Unmarshal(rawBody, &r)
			if err != nil {
				return nil, err
			}
			if r.Code != 0 {
				resp.Err = fmt.Errorf("code: %d, msg: %s", r.Code, r.Msg)
				return nil, nil
			}
			transformedBody, err = json.Marshal(r.Data)
			return
		})

	cli.Container = containers{
		cli: cli.c,
		opt: cli.opt,
	}
	cli.Node = nodes{
		cli: cli.c,
		opt: cli.opt,
	}
	return &cli
}

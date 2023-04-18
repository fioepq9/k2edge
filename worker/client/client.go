package client

import (
	"encoding/json"
	"fmt"
	"k2edge/auth"
	"k2edge/worker/internal/types"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type Client struct {
	opt *ClientOption
	req *req.Client

	Container containerAPI
	Node      nodeAPI
}

type ClientOption struct {
	httpBaseURL string
	secret      string
	token       string
	auther      auth.Auther
}

func (o *ClientOption) HttpBaseURL() string {
	return o.httpBaseURL
}

func (o *ClientOption) WebsocketBaseURL() string {
	base := strings.TrimPrefix(o.httpBaseURL, "http://")
	return fmt.Sprintf("ws://%s", base)
}

type Option func(*ClientOption)

func WithSecret(nodeName, secret string) Option {
	return func(co *ClientOption) {
		co.secret = nodeName + ":" + secret
	}
}

func WithAuther(auther auth.Auther) Option {
	return func(co *ClientOption) {
		co.auther = auther
	}
}

func NewClient(baseurl string, opt ...Option) *Client {
	if !strings.HasPrefix(baseurl, "http://") {
		panic("unsupported protocol")
	}
	var c Client
	c.opt = &ClientOption{
		httpBaseURL: baseurl,
	}
	for _, o := range opt {
		o(c.opt)
	}
	c.req = req.C().
		DevMode().
		SetBaseURL(c.opt.HttpBaseURL()).
		SetCommonHeader("K2edge-Node-Secret", c.opt.secret).
		SetCommonRetryCount(0).
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
			c.opt.token = resp.GetHeader("K2edge-Auth-Token")
			if c.opt.token != "" {
				client.SetCommonHeader("K2edge-Auth-Token", c.opt.token)
				client.Headers.Del("K2edge-Node-Secret")
			} else {
				client.SetCommonHeader("K2edge-Node-Secret", c.opt.secret)
			}
			return nil
		}).
		SetResponseBodyTransformer(func(rawBody []byte, req *req.Request, resp *req.Response) (transformedBody []byte, err error) {
			rawBody = c.opt.auther.Decode(resp.Header.Get("K2edge-Auth-Token"), rawBody)
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

	c.Container = containerAPI{
		req: c.req,
		opt: c.opt,
	}
	c.Node = nodeAPI{
		req: c.req,
		opt: c.opt,
	}
	return &c
}

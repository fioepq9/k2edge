package client

import (
	"encoding/json"
	"fmt"
	"k2edge/master/internal/types"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type Client struct {
	opt *ClientOption
	req *req.Client

	Node      nodeAPI
	Namespace namespaceAPI
	Container containerAPI
	Deployment deploymentAPI
	Job jobAPI
}

type ClientOption struct {
	httpBaseURL string
}

func (o *ClientOption) HttpBaseURL() string {
	return o.httpBaseURL
}

func (o *ClientOption) WebsocketBaseURL() string {
	base := strings.TrimPrefix(o.httpBaseURL, "http://")
	return fmt.Sprintf("ws://%s", base)
}

type Option func(*ClientOption)

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
			return nil
		}).
		SetResponseBodyTransformer(func(rawBody []byte, req *req.Request, resp *req.Response) (transformedBody []byte, err error) {
			var r types.Response
			if resp.StatusCode == 500 {
				return nil, fmt.Errorf("server error, url is %s ", req.URL)
			}
			err = json.Unmarshal(rawBody, &r)
			if err != nil {
				return nil, err
			}
			if r.Code != 0 {
				resp.Err = fmt.Errorf(r.Msg)
				return nil, nil
			}
			transformedBody, err = json.Marshal(r.Data)
			return
		})

	c.Node = nodeAPI{
		req: c.req,
		opt: c.opt,
	}

	c.Namespace = namespaceAPI{
		req: c.req,
		opt: c.opt,
	}

	c.Container = containerAPI{
		req: c.req,
		opt: c.opt,
	}

	c.Deployment = deploymentAPI{
		req: c.req,
		opt: c.opt,
	}

	c.Job = jobAPI{
		req: c.req,
		opt: c.opt,
	}
	return &c
}

package client

import (
	"errors"
	"fmt"
	"k2edge/worker/internal/types"
	"time"

	"github.com/imroc/req/v3"
)

type Client struct {
	*req.Client
}

func NewClient(BaseURL string) *Client {
	var cli Client
	cli.Client = req.C().
		SetBaseURL(BaseURL).
		SetCommonRetryCount(2).
		SetCommonRetryBackoffInterval(time.Second, 5*time.Second).
		AddCommonRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode >= 500
		}).
		WrapRoundTripFunc(func(rt req.RoundTripper) req.RoundTripFunc {
			return func(req *req.Request) (resp *req.Response, err error) {
				resp, err = rt.RoundTrip(req)
				if err != nil {
					return nil, err
				}
				if resp.Err != nil {
					return nil, resp.Err
				}
				if resp.IsErrorState() {
					return nil, fmt.Errorf("bad response, raw content:\n%s", resp.Dump())
				}
				if r, ok := resp.SuccessResult().(types.Response); ok {
					if r.Code != 0 {
						return nil, errors.New(r.Msg)
					}
				}
				return resp, err
			}
		})
	return &cli
}

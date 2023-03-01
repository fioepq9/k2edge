package client

import (
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
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = fmt.Errorf("bad response, raw content:\n%s", resp.Dump())
				return nil
			}
			if r, ok := resp.SuccessResult().(*types.Response); ok {
				if r.Code != 0 {
					resp.Err = fmt.Errorf("code: %d, msg: %s", r.Code, r.Msg)
					return nil
				}
			}
			return nil
		})
	return &cli
}

package client

import (
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

type Client struct {
	*req.Client
}

func NewClient(BaseURL string) *Client {
	var cli Client
	cli.Client = req.C().DevMode().
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
			return nil
		})
	return &cli
}

package client

import (
	"encoding/json"
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
		EnableDumpAll().
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
	return &cli
}

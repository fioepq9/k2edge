package etcdutil

import (
	"context"
	"encoding/json"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetOne[T any](cli *clientv3.Client, ctx context.Context, key string) (result *T, err error) {
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	val := gresp.Kvs[0].Value
	var ret T
	err = json.Unmarshal(val, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func PutOne[T any](cli *clientv3.Client, ctx context.Context, key string, val T) error {
	vbyte, err := json.Marshal(val)
	if err != nil {
		return err
	}
	_, err = cli.KV.Put(ctx, key, string(vbyte))
	if err != nil {
		return err
	}
	return nil
}

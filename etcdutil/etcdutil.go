package etcdutil

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	ErrKeyNotExist error = fmt.Errorf("key not exist")
)

// 获取对应 key 的 value,将得到的 Json 解析并返回
func GetOne[T any](cli *clientv3.Client, ctx context.Context, key string) (result *T, err error) {
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if gresp.Count == 0 {
		return nil, ErrKeyNotExist
	}

	val := gresp.Kvs[0].Value
	var ret T
	err = json.Unmarshal(val, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 将 val 值解析成 Json 格式并添加到对应的 key 中
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

// 将 val 值添加到 key 对应的 []val 中
func AddOne[T any](cli *clientv3.Client, ctx context.Context, key string, val T) error {
	// 获取旧值
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return nil
	}

	if gresp.Count == 0 {
		return ErrKeyNotExist
	}

	var value []T
	err = json.Unmarshal(gresp.Kvs[0].Value, &value)

	if err != nil {
		return err
	}

	// 添加新值
	value = append(value, val)
	vbyte, err := json.Marshal(value)

	if err != nil {
		return err
	}

	//事务提交
	commit, err := cli.Txn(ctx).If(clientv3.Compare(clientv3.ModRevision(key), "=", gresp.Kvs[0].ModRevision)).Then(
		clientv3.OpPut(key, string(vbyte))).Commit()

	if err != nil {
		return err
	}

	if !commit.Succeeded {
		return fmt.Errorf("transaction execution failed when adding %s , please try again", key)
	}
	return nil
}

// 删除 key 下的某个 value 值，通过 lo.filter 来进行过滤, T 不为数组
func DeleteOne[T any](cli *clientv3.Client, ctx context.Context, key string, filter func(item T, index int) bool) error {
	// 获取旧值
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return nil
	}

	if gresp.Count == 0 {
		return ErrKeyNotExist
	}

	var value []T
	err = json.Unmarshal(gresp.Kvs[0].Value, &value)

	if err != nil {
		return err
	}

	// 删除特定值
	value = lo.Filter(value, filter)
	vbyte, err := json.Marshal(value)

	if err != nil {
		return err
	}

	//事务提交
	commit, err := cli.Txn(ctx).If(clientv3.Compare(clientv3.ModRevision(key), "=", gresp.Kvs[0].ModRevision)).Then(
		clientv3.OpPut(key, string(vbyte))).Commit()

	if err != nil {
		return err
	}

	if !commit.Succeeded {
		return fmt.Errorf("transaction execution failed when deleting %s , please try again", key)
	}
	return nil
}

// 根据 metadata 判断资源在是否已经存在
func IsExist(cli *clientv3.Client, ctx context.Context, key string, metadata Metadata) (bool, error) {
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return false, err
	}

	if gresp.Count == 0 {
		return false, ErrKeyNotExist
	}

	var value []Source
	err = json.Unmarshal(gresp.Kvs[0].Value, &value)
	if err != nil {
		return false, err
	}
	
	if err != nil {
		return false, err
	}

	// 判断是否已存在
	for _, v := range value {
		if v.Metadata.Name == metadata.Name && v.Metadata.Kind == metadata.Kind && v.Metadata.Namespace == metadata.Namespace {
			return true, nil
		}
	}
	return false, nil
}

type Metadata struct {
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}

type Source struct {
	Metadata Metadata `json:"metadata"`
}
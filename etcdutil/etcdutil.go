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

// 获取对应 key 的 value,将得到的 Json 解析并返回, T 类型为 []value
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

// 将 val 值添加到 key 对应的 []val 中，T 类型为 value
func AddOne[T any](cli *clientv3.Client, ctx context.Context, key string, val T) error {
	// 获取旧值
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return nil
	}

	if gresp.Count != 0 {
		value := make([]T, 0)
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
	value := []T{val}
	vbyte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = cli.Put(ctx, key, string(vbyte))
	if err != nil {
		return err
	}
	return nil
}

// 删除 key 下的某个 value 值，通过 lo.filter 来进行过滤, T 类型为 value
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

	f := func(item T, idx int) bool {
		return !filter(item, idx)
	}
	// 删除特定值
	value = lo.Filter(value, f)
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

	// 判断是否已存在
	_, found := lo.Find(value, func(item Source) bool {
		return item.Metadata.Kind == metadata.Kind && item.Metadata.Name == metadata.Name && item.Metadata.Namespace == metadata.Namespace
	})
	return found, nil
}

// 判断 namespace 是否存在且可用, namespace 为""则直接返回true
func IsExistNamespace(cli *clientv3.Client, ctx context.Context, namespace string) (bool, error) {
	if namespace == "" {
		return true, nil
	}

	value, err := GetOne[[]Namespace](cli, ctx, "/namespaces")
	if err != nil {
		return false, err
	}

	for _, n := range *value {
		if n.Name == namespace && n.Status == "active" {
			return true, nil
		}
	}
	return false, nil
}

// 根据 nodeName 判断 node 是否存在且可用, 存在就返回
func IsExistNode(cli *clientv3.Client, ctx context.Context, nodeName string) (*Node, bool, error) {
	nodes, err := GetOne[[]Node](cli, ctx, "/nodes")
	if err != nil {
		return nil, false, err
	}

	// 判断结点是否存在
	for _, n := range *nodes {
		if n.Metadata.Name == nodeName && n.Status == "active" {
			ret := new(Node)
			*ret = n
			return ret, true, nil
		}
	}

	// 未找到结点
	return nil, false, fmt.Errorf("cannot find node %s", nodeName)
}

type Metadata struct {
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}

type Source struct {
	Metadata Metadata `json:"metadata"`
}

type Namespace struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	CreateTime int64  `json:"create_time"`
}

type Node struct {
	Metadata     Metadata `json:"metadata"`
	Roles        []string `json:"roles"`
	BaseURL      NodeURL  `json:"base_url"`
	Status       string   `json:"status"`
	RegisterTime int64    `json:"register_time"`
}

type NodeURL struct {
	WorkerURL string `json:"worker_url"`
	MasterURL string `json:"master_url"`
}

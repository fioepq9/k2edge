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
	SystemNamespace = "system"
)

func GenerateKey(kind string, namespace string, name string) string {
	return "/" + kind + "/" + namespace + "/" + name
}

// 获取对应 key 的 value,将得到的 Json 解析并返回, T 类型为 value
func GetOne[T any](cli *clientv3.Client, ctx context.Context, key string) (result *[]T, err error) {
	gresp, err := cli.KV.Get(ctx, key, clientv3.WithPrefix())

	if err != nil {
		return nil, err
	}

	if gresp.Count == 0 {
		return nil, ErrKeyNotExist
	}

	ret := new([]T)
	for _, v := range gresp.Kvs {
		var elem T
		err = json.Unmarshal(v.Value, &elem)
		if err != nil {
			return nil, err
		}

		*ret = append(*ret, elem)
	}
	return ret, nil
}

// 获取对应 key 的值, 且内容为 value[],将得到的 Json 解析并返回, T 类型为 value
func GetArray[T any](cli *clientv3.Client, ctx context.Context, key string) (result *[]T, err error) {
	gresp, err := cli.KV.Get(ctx, key)

	if err != nil {
		return nil, err
	}

	if gresp.Count == 0 {
		return nil, ErrKeyNotExist
	}

	val := gresp.Kvs[0].Value
	var ret []T
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

// 删除某个对应的 key 值
func DeleteOne(cli *clientv3.Client, ctx context.Context, key string) error {
	dresp, err := cli.KV.Delete(ctx, key)

	if err != nil {
		return err
	}

	if dresp.Deleted  == 0 {
		return ErrKeyNotExist
	}

	return nil
}

// 在 key 下添加 value 值，T 类型为 value
func AddOneValue[T any](cli *clientv3.Client, ctx context.Context, key string, val T) error {
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
func DeleteOneValue[T any](cli *clientv3.Client, ctx context.Context, key string, filter func(item T, index int) bool) error {
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

// 根据 metadata 生成的 key 值判断资源在是否已经存在
func IsExistKey(cli *clientv3.Client, ctx context.Context, key string) (bool, error) {
	gresp, err := cli.KV.Get(ctx, key)
	if err != nil {
		return false, err
	}

	return gresp.Count > 0, nil
}

// 判断 namespace 是否存在且可用, namespace 为""则直接返回true
func IsExistNamespace(cli *clientv3.Client, ctx context.Context, namespace string) (bool, error) {
	value, err := GetArray[Namespace](cli, ctx, "/namespaces")
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
	key := GenerateKey("node", SystemNamespace, nodeName)
	
	gresp, err := cli.KV.Get(ctx, key)

	if err != nil {
		return nil, false, err
	}

	// 找不到结点
	if gresp.Count == 0 {
		return nil, false, nil
	}

	var elem Node
	err = json.Unmarshal(gresp.Kvs[0].Value, &elem)
	if err != nil {
		return nil, false, err
	}

	return &elem, true, nil
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

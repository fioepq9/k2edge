package auth

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"k2edge/etcdutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Auther interface {
	GenerateToken(secret string) string
	CheckSecret(secret string) bool
	CheckToken(token string) bool
	Encode(token string, in []byte) (out []byte)
	Decode(token string, in []byte) (out []byte)
}

var _ Auther = (*EtcdAuther)(nil)

type EtcdAuther struct {
	NodeName string
	Node     *Node
	client   *clientv3.Client
}

func NewEtcdAuther(nodeName string, client *clientv3.Client) *EtcdAuther {
	return &EtcdAuther{
		NodeName: nodeName,
		client:   client,
	}
}

// CheckSecret implements Auther
func (a *EtcdAuther) CheckSecret(secret string) bool {
	nameAndSecret := strings.SplitN(secret, ":", 2)
	if len(nameAndSecret) != 2 {
		return false
	}
	name, secret := nameAndSecret[0], nameAndSecret[1]
	nodeKey := etcdutil.GenerateKey("node", "system", name)
	nodes, err := etcdutil.GetOne[Node](
		a.client, context.TODO(), nodeKey,
	)
	if err != nil {
		logx.Debugw("check secret failed", logx.LogField{Key: "error", Value: err})
		return false
	}
	node := &(*nodes)[0]
	return node.Secret == secret
}

// CheckToken implements Auther
func (a *EtcdAuther) CheckToken(token string) bool {
	err := a.GetNode()
	if err != nil {
		logx.Debugw("check token failed", logx.LogField{Key: "error", Value: err})
		return false
	}
	return lo.Contains(a.Node.Status.Token, token)
}

// Decode implements Auther
func (*EtcdAuther) Decode(token string, in []byte) []byte {
	key := []byte(token)[:32]
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(err.Error())
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	out := make([]byte, len(in))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(out, in)
	//去除填充字符串
	out, err = PKCS7UnPadding(out)
	if err != nil {
		return []byte(err.Error())
	}
	return out
}

// Encode implements Auther
func (*EtcdAuther) Encode(token string, in []byte) []byte {
	tokenBytes := []byte(token)[:32]
	block, err := aes.NewCipher(tokenBytes)
	if err != nil {
		return []byte(err.Error())
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	in = PKCS7Padding(in, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, tokenBytes[:blockSize])
	out := make([]byte, len(in))
	//执行加密
	blocMode.CryptBlocks(out, in)
	return out
}

// GenerateToken implements Auther
func (a *EtcdAuther) GenerateToken(secret string) string {
	err := a.GetNode()
	if err != nil {
		logx.Debugw("generate token failed", logx.LogField{Key: "error", Value: err})
		return ""
	}
	token := uuid.NewString()
	a.Node.Status.Token = append(a.Node.Status.Token, token)
	err = a.SetNode()
	if err != nil {
		logx.Debugw("generate token failed", logx.LogField{Key: "error", Value: err})
		return ""
	}
	return token
}

func (a *EtcdAuther) GetNode() error {
	nodeKey := etcdutil.GenerateKey("node", "system", a.NodeName)
	nodes, err := etcdutil.GetOne[Node](
		a.client, context.TODO(), nodeKey,
	)
	if err != nil {
		return err
	}
	a.Node = &(*nodes)[0]
	return nil
}

func (a *EtcdAuther) SetNode() error {
	nodeKey := etcdutil.GenerateKey("node", "system", a.NodeName)
	return etcdutil.PutOne(a.client, context.TODO(), nodeKey, a.Node)
}

// PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充字符串长度
	unpadding := int(origData[length-1])
	//截取切片，删除填充字节，并且返回明文
	return origData[:(length - unpadding)], nil
}

type EncodeResponseWriter struct {
	token  string
	auther Auther
	http.ResponseWriter
}

func NewEncodeResponseWriter(
	token string,
	auther Auther,
	w http.ResponseWriter,
) *EncodeResponseWriter {
	return &EncodeResponseWriter{
		token:          token,
		auther:         auther,
		ResponseWriter: w,
	}
}

func (w *EncodeResponseWriter) Write(in []byte) (int, error) {
	out := w.auther.Encode(w.token, in)
	return w.ResponseWriter.Write(out)
}

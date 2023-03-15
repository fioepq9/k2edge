package logic

import (
	"context"
	"fmt"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttachContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttachContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachContainerLogic {
	return &AttachContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttachContainerLogic) AttachContainer(req *types.AttachContainerRequest) error {
	key := etcdutil.GenerateKey("container", req.Metadata.Namespace, req.Metadata.Name)
	// 判断 container 是否存在, 存在则获取 container 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("container %s does not exist", req.Metadata.Name)
	}

	//根据 container 里 nodeName 去 etcd 里查询的 nodeBaseURL
	containers, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, key)
	if err != nil {
		return err
	}

	container := (*containers)[0]
	// 获取 node 的 BaseURL
	worker, found, err := etcdutil.IsExistNode(l.svcCtx.Etcd, l.ctx, container.ContainerStatus.Node)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("cannot find container %s info", req.Metadata.Name)
	}

	// 向特定的 work 结点发送获取conatiner信息的请求
	cli := client.NewClient(client.WithBaseURL(worker.BaseURL.WorkerURL))
	err = cli.Container.Attach(l.ctx, client.AttachRequest{
		Container: container.ContainerStatus.ContainerID,
		Config: client.AttachConfig{
			Stream:     req.Config.Stream,
			Stdin:      req.Config.Stdin,
			Stdout:     req.Config.Stdout,
			Stderr:     req.Config.Stderr,
			DetachKeys: req.Config.DetachKeys,
			Logs:       req.Config.Logs,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"k2edge/worker/client"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type MigrateContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMigrateContainerLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *MigrateContainerLogic {
	return &MigrateContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MigrateContainerLogic) MigrateContainer(
	req *types.MigrateContainerRequest,
) error {
	ckey := etcdutil.GenerateKey("container", req.Namespace, req.Name)

	// 判断 container 是否存在, 存在则获取 container 信息
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, ckey)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("container %s does not exist", req.Name)
	}
	containers, err := etcdutil.GetOne[types.Container](
		l.svcCtx.Etcd, l.ctx, ckey,
	)
	if err != nil {
		return err
	}
	c := (*containers)[0]

	// 在目标节点创建 container
	targetKey := etcdutil.GenerateKey(
		"node",
		etcdutil.SystemNamespace,
		req.Node,
	)

	targetNodes, err := etcdutil.GetOne[types.Node](
		l.svcCtx.Etcd, l.ctx, targetKey,
	)
	if err != nil {
		return err
	}
	targetNode := (*targetNodes)[0]

	worker, found, err := etcdutil.IsExistNode(
		l.svcCtx.Etcd, l.ctx, targetNode.Metadata.Name,
	)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("cannot find container %s info", req.Name)
	}

	if !worker.Status.Working {
		return errors.New(
			"the node where the container is located is not active",
		)
	}

	targetCli := client.NewClient(worker.BaseURL.WorkerURL)
	targetResp, err := targetCli.Container.Create(
		l.ctx,
		client.CreateContainerRequest{
			ContainerName: req.Name,
			Config: client.ContainerConfig{
				Image:   c.ContainerConfig.Image,
				Command: c.ContainerConfig.Command,
				Args:    c.ContainerConfig.Args,
				Expose: lo.Map(
					c.ContainerConfig.Expose,
					func(item types.ExposedPort, _ int) client.ExposedPort {
						return client.ExposedPort(item)
					},
				),
				Env:   c.ContainerConfig.Env,
				Limit: client.ContainerLimit(c.ContainerConfig.Limit),
			},
		},
	)
	if err != nil {
		return err
	}

	sourceKey := etcdutil.GenerateKey("node", etcdutil.SystemNamespace, c.ContainerStatus.Node)

	sourceNodes, err := etcdutil.GetOne[types.Node](l.svcCtx.Etcd, l.ctx, sourceKey)
	if err != nil {
		return err
	}
	sourceNode := (*sourceNodes)[0]

	worker, found, err = etcdutil.IsExistNode(
		l.svcCtx.Etcd,
		l.ctx,
		sourceNode.Metadata.Name,
	)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("cannot find container %s info", req.Name)
	}

	if !worker.Status.Working {
		return errors.New(
			"the node where the container is located is not active",
		)
	}

	sourceCli := client.NewClient(worker.BaseURL.WorkerURL)
	err = sourceCli.Container.Stop(l.ctx, client.StopContainerRequest{
		ID:      c.ContainerStatus.ContainerID,
		Timeout: 5 * int(time.Minute),
	})
	if err != nil {
		return err
	}
	err = sourceCli.Container.Remove(l.ctx, client.RemoveContainerRequest{
		ID: c.ContainerStatus.ContainerID,
	})
	if err != nil {
		return err
	}

	// 修改 container 的 config 并提交
	c.ContainerStatus.ContainerID = targetResp.ID
	c.ContainerStatus.Node = req.Node
	c.ContainerConfig.NodeName = req.Node
	return etcdutil.PutOne(l.svcCtx.Etcd, l.ctx, ckey, c)
}

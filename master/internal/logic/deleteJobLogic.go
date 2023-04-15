package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJobLogic {
	return &DeleteJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteJobLogic) DeleteJob(req *types.DeleteJobRequest) error {
	key := etcdutil.GenerateKey("job", req.Namespace, req.Name)
	found, err := etcdutil.IsExistKey(l.svcCtx.Etcd, l.ctx, key)

	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("job %s does not exists", req.Name)
	}

	// 获取Job相关的所有container
	keyc := "/container"

	containers, err := etcdutil.GetOne[types.Container](l.svcCtx.Etcd, l.ctx, keyc)
	if err != nil {
		if errors.Is(err, etcdutil.ErrKeyNotExist) {
			return nil
		}
		return err
	}

	logic := NewDeleteContainerLogic(l.ctx, l.svcCtx)
	for _, container := range *containers {
		if container.ContainerConfig.Job == req.Namespace+"/"+req.Name {
			err = logic.DeleteContainer(&types.DeleteContainerRequest{
				Namespace: container.Metadata.Namespace,
				Name:      container.Metadata.Name,
			})
			if err != nil {
				if !strings.Contains(err.Error(), "No such container:") {
					return err
				}

				err = etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key)
				if err != nil {
					return err
				}
			}
		}
	}
	return etcdutil.DeleteOne(l.svcCtx.Etcd, l.ctx, key)
}

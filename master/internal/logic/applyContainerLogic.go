package logic

import (
	"context"
	"fmt"

	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyContainerLogic {
	return &ApplyContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyContainerLogic) ApplyContainer(req *types.ApplyContainerRequest) error {
	// 获取容器信息以便恢复
	getFunc := NewGetContainerLogic(l.ctx, l.svcCtx)
	config, err := getFunc.GetContainer(&types.GetContainerRequest{
		Namespace: req.Container.Metadata.Namespace,
		Name: req.Container.Metadata.Name,
	})

	if err != nil {
		return fmt.Errorf("backup container config: %s", err)
	}

	// 删除容器
	deleteFunc := NewDeleteContainerLogic(l.ctx, l.svcCtx)
	err = deleteFunc.DeleteContainer(&types.DeleteContainerRequest{
		Namespace: req.Container.Metadata.Namespace,
		Name: req.Container.Metadata.Name,
	})

	if err != nil {
		return fmt.Errorf("delete container: %s", err)
	}

	// 创建容器
	createFunc := NewCreateContainerLogic(l.ctx, l.svcCtx)
	err = createFunc.CreateContainer(&types.CreateContainerRequest{
		Container: req.Container,
	})

	if err != nil {
		errc := createFunc.CreateContainer(&types.CreateContainerRequest{
			Container: config.Container,
		})
		if errc != nil {
			return fmt.Errorf("recover container: %s", errc)
		}

		return fmt.Errorf("create container: %s", err)
	}

	return nil
}

package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	typesInternal "k2edge/worker/internal/types"

	"github.com/docker/docker/api/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttachLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttachLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachLogic {
	return &AttachLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttachLogic) Attach(req *typesInternal.AttachRequest) error {
	_, err := l.svcCtx.DockerClient.ContainerAttach(l.ctx, req.Container, types.ContainerAttachOptions{
		Stream:     req.Config.Stream,
		Stdin:      req.Config.Stdin,
		Stdout:     req.Config.Stdout,
		Stderr:     req.Config.Stderr,
		DetachKeys: req.Config.DetachKeys,
		Logs:       req.Config.Logs,
	})
	if err != nil {
		return err
	}
	// todo: websocket
	return nil
}

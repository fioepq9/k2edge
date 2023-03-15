package logic

import (
	"context"
	"io"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
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

func (l *AttachLogic) Attach(req *types.AttachRequest) (io.ReadWriteCloser, error) {
	d := l.svcCtx.Docker
	stream, err := d.ContainerAttach(l.ctx, req.Container, dtypes.ContainerAttachOptions{
		Stream:     req.Stream,
		Stdin:      req.Stdin,
		Stdout:     req.Stdout,
		Stderr:     req.Stderr,
		DetachKeys: req.DetachKeys,
		Logs:       req.Logs,
	})
	if err != nil {
		return nil, err
	}
	return stream.Conn, nil
}

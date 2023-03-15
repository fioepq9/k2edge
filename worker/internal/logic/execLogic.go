package logic

import (
	"bufio"
	"context"
	"io"

	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	dtypes "github.com/docker/docker/api/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecLogic {
	return &ExecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecLogic) Exec(req *types.ExecRequest) (io.ReadWriteCloser, error) {
	d := l.svcCtx.Docker
	cresp, err := d.ContainerExecCreate(l.ctx, req.Container, dtypes.ExecConfig{
		User:         req.User,
		Privileged:   req.Privileged,
		Tty:          req.Tty,
		AttachStdin:  req.AttachStdin,
		AttachStderr: req.AttachStderr,
		AttachStdout: req.AttachStdout,
		Detach:       req.Detach,
		DetachKeys:   req.DetachKeys,
		Env:          req.Env,
		WorkingDir:   req.WorkingDir,
		Cmd:          req.Cmd,
	})
	if err != nil {
		return nil, err
	}
	aresp, err := d.ContainerExecAttach(l.ctx, cresp.ID, dtypes.ExecStartCheck{
		Detach: req.Detach,
		Tty:    req.Tty,
	})
	if err != nil {
		return nil, err
	}
	if req.Tty {
		return aresp.Conn, nil
	}
	return nopCloser{ReadWriter: bufio.NewReadWriter(aresp.Reader, nil)}, nil
}

type nopCloser struct {
	io.ReadWriter
}

func (n nopCloser) Close() error {
	return nil
}

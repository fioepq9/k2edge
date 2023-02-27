package logic

import (
	"context"

	"k2edge/worker/internal/svc"
	typesInternal "k2edge/worker/internal/types"

	"github.com/docker/docker/api/types"
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

func (l *ExecLogic) Exec(req *typesInternal.ExecRequest) error {
	_, err := l.svcCtx.DockerClient.ContainerExecCreate(l.ctx, req.Container, types.ExecConfig{
		User:         req.Config.User,
		Privileged:   req.Config.Privileged,
		Tty:          req.Config.Tty,
		AttachStdin:  req.Config.AttachStdin,
		AttachStderr: req.Config.AttachStderr,
		AttachStdout: req.Config.AttachStdout,
		Detach:       req.Config.Detach,
		DetachKeys:   req.Config.DetachKeys,
		Env:          req.Config.Env,
		WorkingDir:   req.Config.WorkingDir,
		Cmd:          req.Config.Cmd,
	})
	if err != nil {
		return err
	}
	return nil
}

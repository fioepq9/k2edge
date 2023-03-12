package logic

import (
	"context"
	"io"
	"net"
	"sync"

	"k2edge/worker/internal/svc"
	typesInternal "k2edge/worker/internal/types"

	"github.com/docker/docker/api/types"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ws     *websocket.Conn
}

func NewExecLogic(ctx context.Context, svcCtx *svc.ServiceContext, ws *websocket.Conn) *ExecLogic {
	return &ExecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		ws:     ws,
	}
}

func (l *ExecLogic) Exec(req *typesInternal.ExecRequest) error {
	d := l.svcCtx.DockerClient
	cresp, err := d.ContainerExecCreate(l.ctx, req.Container, types.ExecConfig{
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
	aresp, err := d.ContainerExecAttach(l.ctx, cresp.ID, types.ExecStartCheck{
		Detach: req.Config.Detach,
		Tty:    req.Config.Tty,
	})
	if err != nil {
		return err
	}
	if !req.Config.Tty {
		resp, err := io.ReadAll(aresp.Reader)
		if err != nil {
			return err
		}
		l.ws.WriteMessage(websocket.BinaryMessage, resp)
		l.ws.WriteMessage(websocket.CloseMessage, []byte("finish"))
		return nil
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go l.Write(aresp.Conn, &wg)
	go l.Read(aresp.Conn, &wg)
	wg.Wait()
	return nil
}

func (l *ExecLogic) Read(conn net.Conn, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		wt, err := l.ws.NextWriter(websocket.BinaryMessage)
		if err != nil {
			return err
		}
		if _, err = io.Copy(wt, conn); err != nil {
			return err
		}
	}
}

func (l *ExecLogic) Write(conn net.Conn, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		_, rd, err := l.ws.NextReader()
		if err != nil {
			return err
		}
		if _, err = io.Copy(conn, rd); err != nil {
			return err
		}
	}
}

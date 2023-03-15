package handler

import (
	"io"
	"net/http"
	"sync"

	"k2edge/worker/internal/logic"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AttachHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttachRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		ws, err := svcCtx.Websocket.Upgrade(w, r, nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer ws.Close()
		l := logic.NewAttachLogic(r.Context(), svcCtx)
		rw, err := l.Attach(&req)
		if err != nil {
			msg := websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, err.Error())
			ws.WriteMessage(websocket.CloseMessage, msg)
			return
		}
		defer rw.Close()
		session := AttachSession{
			ws:     ws,
			stream: rw,
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			for session.Read() == nil {
			}
		}()
		go func() {
			defer wg.Done()
			for session.Write() == nil {
			}
		}()
		wg.Wait()
		msg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "close")
		ws.WriteMessage(websocket.CloseMessage, msg)
	}
}

type AttachSession struct {
	ws     *websocket.Conn
	stream io.ReadWriteCloser
}

func (s *AttachSession) Read() error {
	_, r, err := s.ws.NextReader()
	if err != nil {
		return err
	}
	_, err = io.Copy(s.stream, r)
	if err != nil {
		return err
	}
	return nil
}

func (s *AttachSession) Write() error {
	w, err := s.ws.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	_, err = io.CopyN(w, s.stream, 1)
	if err != nil {
		return err
	}
	return nil
}

package handler

import (
	"io"
	"net/http"

	"k2edge/worker/internal/logic"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LogsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		ws, err := svcCtx.Websocket.Upgrade(w, r, nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		defer ws.Close()

		l := logic.NewLogsLogic(r.Context(), svcCtx)
		rd, err := l.Logs(&req)
		if err != nil {
			msg := websocket.FormatCloseMessage(
				websocket.CloseAbnormalClosure,
				err.Error(),
			)
			ws.WriteMessage(websocket.CloseMessage, msg)
			return
		}
		defer rd.Close()
		session := LogsSession{
			ws: ws,
			rd: rd,
		}
		for session.Write() == nil {
		}
		msg := websocket.FormatCloseMessage(
			websocket.CloseNormalClosure,
			"close",
		)
		ws.WriteMessage(websocket.CloseMessage, msg)
	}
}

type LogsSession struct {
	ws *websocket.Conn
	rd io.ReadCloser
}

func (s *LogsSession) Write() error {
	w, err := s.ws.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	_, err = io.CopyN(w, s.rd, 1)
	if err != nil {
		return err
	}
	return nil
}

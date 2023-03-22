package handler

import (
	"net/http"
	"sync"

	"k2edge/worker/internal/logic"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExecHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExecRequest
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
		l := logic.NewExecLogic(r.Context(), svcCtx)
		rw, err := l.Exec(&req)
		if err != nil {
			ws.WriteMessage(websocket.TextMessage, []byte(err.Error() + "\n"))
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

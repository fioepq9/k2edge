package handler

import (
	"net/http"
	"sync"

	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AttachContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttachContainerRequest
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
		l := logic.NewAttachContainerLogic(r.Context(), svcCtx)
		rw, err := l.AttachContainer(&req)
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

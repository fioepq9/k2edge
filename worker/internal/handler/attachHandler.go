package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k2edge/worker/internal/logic"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"
)

func AttachHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttachRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAttachLogic(r.Context(), svcCtx)
		resp, err := l.Attach(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
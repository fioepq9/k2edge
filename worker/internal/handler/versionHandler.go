package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k2edge/worker/internal/logic"
	"k2edge/worker/internal/svc"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

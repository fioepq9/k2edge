package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
)

func DeleteCronJobHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCronJobRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteCronJobLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCronJob(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

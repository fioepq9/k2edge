package handler

import (
	"net/http"

	"k2edge/worker/internal/logic"
	"k2edge/worker/internal/svc"
	"k2edge/worker/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateContainerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateContainerLogic(r.Context(), svcCtx)
		resp, err := l.CreateContainer(&req)
		var body types.Response
		if err != nil {
			body.Code = -1
			body.Msg = err.Error()
		} else {
			body.Msg = "success"
			body.Data = resp
		}
		httpx.OkJsonCtx(r.Context(), w, body)
	}
}

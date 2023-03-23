package handler

import (
	"net/http"

	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HostTopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHostTopLogic(r.Context(), svcCtx)
		resp, err := l.HostTop()
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

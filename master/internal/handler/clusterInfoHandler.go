package handler

import (
	"net/http"

	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClusterInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewClusterInfoLogic(r.Context(), svcCtx)
		resp, err := l.ClusterInfo()
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

package handler

import (
	"net/http"

	"k2edge/master/internal/types"
	"k2edge/master/internal/logic"
	"k2edge/master/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListNamespaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListNamespaceLogic(r.Context(), svcCtx)
		resp, err := l.ListNamespace()
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

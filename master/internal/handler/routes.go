// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"k2edge/master/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateCronJobHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetCronJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteCronJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply",
				Handler: ApplyCronJobHandler(serverCtx),
			},
		},
		rest.WithPrefix("/cronjob"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateNamespaceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetNamespaceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteNamespaceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply",
				Handler: ApplyNamespaceHandler(serverCtx),
			},
		},
		rest.WithPrefix("/namespace"),
	)
}

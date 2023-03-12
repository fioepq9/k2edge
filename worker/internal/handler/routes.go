// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"k2edge/worker/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/version",
				Handler: VersionHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/container/create",
					Handler: CreateContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/container/remove",
					Handler: RemoveContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/container/stop",
					Handler: StopContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/container/start",
					Handler: StartContainerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/container/status",
					Handler: ContainerStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/container/list",
					Handler: ListContainersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/container/exec",
					Handler: ExecHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/container/attach",
					Handler: AttachHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/node/top",
					Handler: NodeTopHandler(serverCtx),
				},
			}...,
		),
	)
}

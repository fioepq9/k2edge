// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"k2edge/master/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: ClusterInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/cluster"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: ListContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply",
				Handler: ApplyContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/attach",
				Handler: AttachContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/exec",
				Handler: ExecContainerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/logs",
				Handler: LogsContainerHandler(serverCtx),
			},
		},
		rest.WithPrefix("/container"),
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
			{
				Method:  http.MethodGet,
				Path:    "/rollout/history",
				Handler: HistoryCronJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/rollout/undo",
				Handler: UndoCronJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logs",
				Handler: LogsCronJobHandler(serverCtx),
			},
		},
		rest.WithPrefix("/cronjob"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply",
				Handler: ApplyDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/rollout/history",
				Handler: HistoryDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/rollout/undo",
				Handler: UndoDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/scale",
				Handler: ScaleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/attach",
				Handler: AttachDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exec",
				Handler: ExecDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logs",
				Handler: LogsDeploymentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/deployment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateJobHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply",
				Handler: ApplyJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logs",
				Handler: LogsJobHandler(serverCtx),
			},
		},
		rest.WithPrefix("/job"),
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
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: ListNamespaceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteNamespaceHandler(serverCtx),
			},
		},
		rest.WithPrefix("/namespace"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: registerNodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/top",
				Handler: NodeTopHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cordon",
				Handler: CordonHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/uncordon",
				Handler: UncordonHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/drain",
				Handler: DrainHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteNodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/hostTop",
				Handler: HostTopHandler(serverCtx),
			},
		},
		rest.WithPrefix("/node"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply",
				Handler: ApplyTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/token"),
	)
}

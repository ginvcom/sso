// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"sso/service/auth/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sign-in",
				Handler: signInHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/verify-request",
				Handler: verifyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/session-menus",
				Handler: menusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-out",
				Handler: signOutHandler(serverCtx),
			},
		},
	)
}

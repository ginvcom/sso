// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	object "sso/ssoms/api/internal/handler/object"
	permission "sso/ssoms/api/internal/handler/permission"
	role "sso/ssoms/api/internal/handler/role"
	user "sso/ssoms/api/internal/handler/user"
	"sso/ssoms/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: user.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: user.AddUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: user.UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: user.UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: user.DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/filterOptions",
				Handler: user.UserFilterOptionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/assignedRoles",
				Handler: user.AssignedRolesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/assignRole",
				Handler: user.AssignRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/permissions",
				Handler: user.UserPermissionsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: role.RoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: role.AddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: role.RoleDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: role.UpdateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: role.DeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/options",
				Handler: role.RoleOptionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/assignedUsers",
				Handler: role.AssignedUsersHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/assignUser",
				Handler: role.AssignUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/deassignUser",
				Handler: role.DeassignUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/inheritances",
				Handler: role.InheritancesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addInheritance",
				Handler: role.ExtendRoleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/role"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: object.ObjectListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: object.AddObjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/detail",
				Handler: object.ObjectDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: object.UpdateObjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: object.DeleteObjectHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/roleOperations",
				Handler: object.RoleOperationsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/object"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/RolePermissions",
				Handler: permission.RolePermissionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/grant",
				Handler: permission.GrantHandler(serverCtx),
			},
		},
		rest.WithPrefix("/permission"),
	)
}

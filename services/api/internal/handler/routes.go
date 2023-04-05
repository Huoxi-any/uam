// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	group "uam/services/api/internal/handler/group"
	permission "uam/services/api/internal/handler/permission"
	role "uam/services/api/internal/handler/role"
	user "uam/services/api/internal/handler/user"
	"uam/services/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ApiAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/permission-ids",
					Handler: user.UserPermIdsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/permission-keys",
					Handler: user.UserPermKeysHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user-group",
					Handler: user.UserGroupsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user-group",
					Handler: user.AddUserGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user-group",
					Handler: user.RemoveUserGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user-role",
					Handler: user.AddUserRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user-role",
					Handler: user.RemoveUserRoleHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/uam/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ApiAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/permissions",
					Handler: permission.PermListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/permission/batch",
					Handler: permission.AddPermBatchHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/permission",
					Handler: permission.DeletePermHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/uam/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ApiAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/roles",
					Handler: role.RoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role",
					Handler: role.GetRoleByNameHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.UpsertRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/role",
					Handler: role.UpdateRolePermHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/uam/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ApiAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/groups",
					Handler: group.GroupListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/group",
					Handler: group.GetGroupByNameHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/group",
					Handler: group.UpsertGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/group/users",
					Handler: group.GetUserByGroupNameHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/group/permission-ids",
					Handler: group.GroupPermIdsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/group/permission-keys",
					Handler: group.GroupPermKeysHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/uam/api/v1"),
	)
}

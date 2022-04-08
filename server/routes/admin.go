package routes

import (
	"gf-admin/app/handler/admin/auth"
	"gf-admin/app/handler/admin/permission"
	"gf-admin/app/handler/admin/role"
	"gf-admin/app/handler/admin/scaffold"
	"gf-admin/app/handler/admin/user"
	"gf-admin/app/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/api", func(api *ghttp.RouterGroup) {
		api.Middleware(middleware.Response, ghttp.MiddlewareCORS)
		api.Group("/v1", func(v1 *ghttp.RouterGroup) {
			v1.Bind(new(scaffold.Scaffold))
			v1.Bind(new(user.Api))
			v1.Bind(new(auth.Auth))
			v1.Bind(new(role.Api))
			v1.Bind(new(permission.Api))
		})
	})
}

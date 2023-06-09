package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"myblog-gf/internal/controller/manage"
	"myblog-gf/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 获得一个默认的Server对象
			s := g.Server()
			// 允许路由覆盖
			s.SetRouteOverWrite(true)
			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(middleware.JsonResponse)
			//	group.Bind(
			//		controller.Hello,
			//	)
			//})
			s.Group("/manage", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.CrossDomain,
					middleware.JsonResponse,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.ManageAuth)
					group.Bind(
						manage.LoginController,
						manage.CategoryController,
						manage.TagController,
						manage.ArticleController,
					)
				})

				// 鉴权例外
				group.Map(g.Map{
					"POST:/login": manage.LoginController.Login,
				})
			})
			// 设置端口
			s.SetPort(9090)
			// 执行Server的监听运行
			s.Run()
			return nil
		},
	}
)

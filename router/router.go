package router

import (
	"myproject/app/api"

	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Use(func(r *ghttp.Request) {
		r.Response.CORSDefault()
		r.Middleware.Next()
	})

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)
	})
}

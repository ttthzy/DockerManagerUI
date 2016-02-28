package routes

import (
	s "github.com/ttthzy/DockerManagerUI/services"
	"gopkg.in/macaron.v1"
)

func (r *RouteHandel) RouteMap() {
	m := r.m
	/// 页面模板处理
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["title"] = "Docker镜像管理"
		ctx.HTML(200, "docker/images")
	})

	/// 账号服务组
	m.Group("/passwport", func() {
		m.Get("/login.html", func(ctx *macaron.Context) {
			ctx.Data["title"] = "用户登录"
			ctx.HTML(200, "passwport/login")
		})
	})

	/// Docker API处理组
	h := new(s.BaseHandle)
	m.Group("/docker", func() {
		m.Get("/getimages", h.GetImages)
		//m.Any("/containers", h.GetImages)
	})

	/// 启动 httpserver 服务
	r.HttpServerRun()
}

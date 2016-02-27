package main

import (
	"log"
	"net/http"

	s "github.com/ttthzy/DockerManagerUI/services"
	"gopkg.in/macaron.v1"
)

func main() {

	/// macaron初始化
	m := macaron.Classic()
	m.Use(macaron.Renderer())                                                     // 允许服务器渲染模板
	m.Use(macaron.Static("static/file/"))                                         // 设置静态目录
	m.Use(macaron.Renderer(macaron.RenderOptions{Directory: "static/templates"})) // 模板文件目录，默认为 "templates"

	/// 页面模板处理
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["title"] = "Docker镜像管理"
		ctx.HTML(200, "docker/images")
	})

	m.Group("/passwport", func() {
		m.Get("/login.html", func(ctx *macaron.Context) {
			ctx.Data["title"] = "用户登录"
			ctx.HTML(200, "passwport/login")
		})
	})

	/// Docker API处理
	h := new(s.BaseHandle)
	m.Group("/docker", func() {
		m.Get("/getimages", h.GetImages)
		//m.Any("/containers", h.GetImages)
	})

	/// http服务启动
	log.Println("Server is running...")
	http.ListenAndServe("0.0.0.0:8080", m)
	//http.ListenAndServeTLS(":8080", "server.crt", "server.key", m)
}

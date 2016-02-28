package routes

import (
	api "github.com/ttthzy/mydockerui/services/api"
	controller "github.com/ttthzy/mydockerui/services/controller"
)

func (r *RouteHandel) RouteMap() {

	/////////////////////////////////////////////////////
	/* --------------- html模板处理  ----------------- */
	chandle := new(controller.BaseHandle)

	/* docker html */
	r.m.Group("/docker", func() {
		r.m.Get("/images.html", chandle.DockerImages) //镜像列表页
	})

	/* passport html */
	r.m.Group("/passwport", func() {
		r.m.Get("/login.html", chandle.UserLogin) //用户登录页
	})
	/////////////////////////////////////////////////////

	/////////////////////////////////////////////////////
	/* --------------- webapi处理  ------------------- */
	ahandle := new(api.BaseHandle)

	/* docker remote api */
	r.m.Group("/docker", func() {
		r.m.Get("/getimages", ahandle.GetImages)
		//m.Any("/containers", h.GetImages)
	})
	/////////////////////////////////////////////////////

	/////////////////////////////////////////////////////
	/* ---------------- 启动httpserver --------------- */
	r.HttpServerRun()
	/////////////////////////////////////////////////////
}

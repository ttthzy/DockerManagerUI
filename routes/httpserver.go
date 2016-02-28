package routes

import "log"

func (r *RouteHandel) HttpServerRun() {
	/// http服务启动
	log.Println("Server is running...")
	r.m.Run(8080)
}

package routes

import (
	"log"
	"runtime"
)

func (r *RouteHandel) HttpServerRun() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/// http服务启动
	log.Println("Server is running...")
	r.m.Run(8080)
}

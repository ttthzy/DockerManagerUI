package routes

import (
	"log"
	"net/http"
)

func (r *RouteHandel) HttpServerRun() {
	/// http服务启动
	log.Println("Server is running...")
	http.ListenAndServe("0.0.0.0:8080", r.m)
	//http.ListenAndServeTLS(":8080", "server.crt", "server.key", m)
}

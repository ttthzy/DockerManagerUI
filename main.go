package main

import (
	"log"
	"net/http"
	"gopkg.in/macaron.v1"
	"github.com/ttthzy/DockerManagerUI/Services"
)

func main() {
	m := macaron.Classic()
	m.Get("/", myHandler)
	m.Get("/getimage", Services.GetImages)

	log.Println("Server is running...")
	http.ListenAndServe("0.0.0.0:8080", m)
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}
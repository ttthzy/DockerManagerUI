package main

import (
	"log"
	"net/http"
	"github.com/ttthzy/DockerManagerUI/Services"
	"gopkg.in/macaron.v1"
)

type Person struct {
    Name string
    Age  int
    Sex  string
}

func main() {
	m := macaron.Classic()
    m.Use(macaron.Renderer())
    
    
    // m.Get("/xml", func(ctx *macaron.Context) {
    //     p := Person{"Unknwon", 21, "male"}
    //     ctx.XML(200, &p)
    // })
    // m.Get("/json", func(ctx *macaron.Context) {
    //     p := Person{"Unknwon", 21, "male"}
    //     ctx.JSON(200, &p)
    // })
    // m.Get("/raw", func(ctx *macaron.Context) {
    //     ctx.RawData(200, []byte("raw data goes here"))
    // })
    // m.Get("/text", func(ctx *macaron.Context) {
    //     ctx.PlainText(200, []byte("plain text goes here"))
    // })
    
    
    

	m.Get("/", func(ctx *macaron.Context) {
        ctx.Data["Name"] = "jeremy"
        ctx.HTML(200, "docker/hello") // 200 为响应码
    })

	m.Get("/getimage", Services.GetImages)

	log.Println("Server is running...")
	http.ListenAndServe("0.0.0.0:8080", m)
    //http.ListenAndServeTLS(":8080", "server.crt", "server.key", m)
}

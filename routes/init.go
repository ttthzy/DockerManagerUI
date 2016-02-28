package routes

import (
	"github.com/go-macaron/cache"
	_ "github.com/go-macaron/cache/redis"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	_ "github.com/go-macaron/session/redis"
	"gopkg.in/macaron.v1"
)

type RouteHandel struct {
	m *macaron.Macaron
}

func (r *RouteHandel) MaInit() {
	/// macaron初始化
	r.m = macaron.Classic()
	r.m.Use(macaron.Renderer())                                                     // 允许服务器渲染模板
	r.m.Use(macaron.Static("static/file/"))                                         // 设置静态目录
	r.m.Use(macaron.Renderer(macaron.RenderOptions{Directory: "static/templates"})) // 模板文件目录，默认为 "templates"
	r.m.Use(macaron.Logger())                                                       //开启路由日志
	r.m.Use(session.Sessioner())                                                    //启用session
	r.m.Use(csrf.Csrfer())                                                          // 开启 CSRF防御机制
	/// 启用缓存服务
	r.m.Use(cache.Cacher(cache.Options{
		Adapter:       "redis",
		AdapterConfig: "addr=192.168.108.128:32768,password=secondlife",
		OccupyMode:    false, //true则独占redis
	}))

	r.RouteMap()
}

package controller

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

// Docker镜像列表页
func (h *BaseHandle) DockerImages(ctx *macaron.Context, sess session.Store) {

	if sess.Get("uid") == nil {
		ctx.Redirect("/passwport/login.html")
		return
	}

	ctx.Data["title"] = "Docker镜像管理"
	ctx.HTML(200, "docker/images")
}

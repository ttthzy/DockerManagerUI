package controller

import "gopkg.in/macaron.v1"

// Docker镜像列表页
func (h *BaseHandle) DockerImages(ctx *macaron.Context) {
	ctx.Data["title"] = "Docker镜像管理"
	ctx.HTML(200, "docker/images")
}

// 用户登录
func (h *BaseHandle) UserLogin(ctx *macaron.Context) {
	ctx.Data["title"] = "用户登录"
	ctx.HTML(200, "passport/login")
}

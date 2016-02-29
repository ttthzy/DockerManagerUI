package api

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func (h *BaseHandle) CheckLogin(ctx *macaron.Context, sess session.Store) {

	username := ctx.Query("username")
	password := ctx.Query("password")

	if username == "admin" && password == "admin" {
		sess.Set("uid", username)
		ctx.Redirect("/docker/images.html")
	} else {
		sess.Set("error", "账号或密码不正确！")
		ctx.Redirect("/passwport/login.html")
	}

}

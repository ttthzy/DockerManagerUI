package controller

import (
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

// 用户登录
func (h *BaseHandle) UserLogin(ctx *macaron.Context, sess session.Store, x csrf.CSRF) {

	ctx.Data["error"] = sess.Get("error")

	ctx.Data["csrf_token"] = x.GetToken()
	ctx.Data["title"] = "用户登录"
	ctx.HTML(200, "passport/login")
}

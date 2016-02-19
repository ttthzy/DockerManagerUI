package Services

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func GetImages(ctx *macaron.Context) {
	ctx.Resp.WriteHeader(200) // HTTP 200
	ctx.Resp.Write([]byte("the request path is: " + ctx.Req.RequestURI))
	// ... 处理内容
}


// 验证一个 API 密钥
func GetReg(ctx *macaron.Context) {
	if ctx.Req.Header.Get("X-API-KEY") != "secret123" {
		ctx.Resp.WriteHeader(http.StatusUnauthorized)
	}
}
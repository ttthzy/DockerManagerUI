package api

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/macaron.v1"
)

func (h *BaseHandle) GetImages(ctx *macaron.Context) {

	message := "ok"
	furl := "/images/json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", DockerSvrAddress+furl, nil)
	if err != nil {
		message = "接口错误"
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		message = "发送失败"
	}

	h.Result(string(body), 200, message, ctx.Resp)

}

// 验证一个 API 密钥
func (h *BaseHandle) GetReg(ctx *macaron.Context) {
	if ctx.Req.Header.Get("X-API-KEY") != "secret123" {
		ctx.Resp.WriteHeader(http.StatusUnauthorized)
	}
}

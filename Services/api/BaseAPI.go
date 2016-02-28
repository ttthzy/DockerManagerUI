package api

import (
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

const (
	DockerSvrAddress = "http://192.168.108.128:2375"
)

/// http请求接口父类
type BaseHandle struct {
	returnJson baseJsonBean
}

/// API返回的json约定格式
type baseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

/// 实例化一个 BaseJsonBean
func (h *BaseHandle) NewBaseJsonBean() *baseJsonBean {
	return &h.returnJson
}

/// 返回服务器公共约定的json
func (h *BaseHandle) Result(data interface{}, code int, message string, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json;charset=utf-8") //返回数据格式是json
	w.WriteHeader(code)

	result := h.NewBaseJsonBean()
	result.Data = data
	result.Code = code
	result.Message = message

	bytes, _ := ffjson.Marshal(result)

	w.Write(bytes)
}

/// 允许跨域访问
func (h *BaseHandle) AllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问的域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

}

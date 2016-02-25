package Services

import (
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func RetrunResult(data interface{}, code int, message string, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json;charset=utf-8") //返回数据格式是json
	w.WriteHeader(code)

	result := NewBaseJsonBean()
	result.Data = data
	result.Code = code
	result.Message = message

	bytes, _ := ffjson.Marshal(result)

	w.Write(bytes)
}

///设置跨域访问
func setAllowOrigin(w http.ResponseWriter, domain string) {
	w.Header().Set("Access-Control-Allow-Origin", domain)          //允许访问的域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

}

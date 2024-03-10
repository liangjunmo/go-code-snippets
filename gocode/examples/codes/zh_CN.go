package codes

import (
	"github.com/liangjunmo/goutil/gocode"
)

var zhCn = map[gocode.Code]string{
	OK:                  "OK",
	Unknown:             "未知错误",
	Timeout:             "请求超时",
	NotFound:            "资源不存在",
	InvalidRequest:      "请求错误",
	InternalServerError: "服务端错误",
}

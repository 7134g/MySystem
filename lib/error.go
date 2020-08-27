package lib

import (
	"MySystem/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

// 通用错误处理
func Err(code int, msg string, err error) Response {
	response := Response{
		Code: code,
		Msg:  msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		response.Error = err.Error()
	}

	return response
}

// 参数异常
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeDBError, msg, err)
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ParamErr("JSON类型不匹配", err)
	}

	return ParamErr("参数错误", err)
}

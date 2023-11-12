package errcode

import (
	"fmt"
	"net/http"
)

// Errors 常用的一些错误处理公共方法，标准化错误输出

type Errors struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Errors {
	// 创建新的 Error 实例之前，先进行查重的校验
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Errors{
		code:    code,
		msg:     msg,
		details: []string{},
	}
}

func (e *Errors) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

func (e *Errors) Code() int {
	return e.code
}

func (e *Errors) Msg() string {
	return e.msg
}

func (e *Errors) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Errors) Details() []string {
	return e.details
}

func (e *Errors) WithDetails(details ...string) *Errors {
	newError := *e
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// StatusCode
// 针对一些特定错误码进行状态码的转换，因为不同的内部错误码在 HTTP 状态码中都代表着不同的意义，我们需要将其区分开来，
// 便于客户端以及监控/报警等系统的识别和监听。
func (e *Errors) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case ErrorUserExit.Code():
		return http.StatusOK
	case ErrorUserPassword.Code():
		return http.StatusOK
	case ErrorUserNotExit.Code():
		return http.StatusOK
	}
	return http.StatusInternalServerError
}

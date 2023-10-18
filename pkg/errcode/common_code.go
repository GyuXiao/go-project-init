package errcode

// 公共错误码

var (
	Success                   = NewError(2000000, "成功")
	ServerError               = NewError(1000000, "服务内部错误")
	InvalidParams             = NewError(1000001, "入参错误")
	NotFound                  = NewError(1000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(1000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(1000005, "鉴权失败， Token 超时")
	UnauthorizedTokenGenerate = NewError(1000006, "鉴权失败， Token 生成失败")
	TooManyRequests           = NewError(1000007, "请求过多")
	InvalidToken              = NewError(1000008, "Token 无效")
)

// Package errcode
/**
  @author: zk.xiao
  @date: 2022/6/1
  @note:
**/
package errcode

// User 错误码

var (
	ErrorUserSignupFail = NewError(20040001, "用户注册失败")
	ErrorUserExit       = NewError(20040002, "用户已经存在")
	ErrorUserNotExit    = NewError(20040003, "用户不存在")
	ErrorUserPassword   = NewError(20040004, "用户密码错误")
	ErrorUserLoginFail  = NewError(20040005, "用户登陆失败")
)

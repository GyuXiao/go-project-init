package v2

import (
	"GyuBlog/global"
	"GyuBlog/pkg/app"
	"GyuBlog/pkg/errcode"
	"GyuBlog/service"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) SignupHandler(c *gin.Context) {
	// 参数校验
	param := service.UserSignupRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 业务处理——注册用户
	svc := service.New(c.Request.Context())
	err := svc.Signup(&param)
	if err != nil {
		if err == errcode.ErrorUserExit {
			// 用户已经存在
			response.ToErrorResponse(errcode.ErrorUserExit)
			return
		}
		global.Logger.Errorf(c, "svc.Signup failed, err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserSignupFail.WithDetails(errs.Errors()...))
		return
	}

	// 业务响应
	response.ToErrorResponse(errcode.Success.WithDetails([]string{param.UserName, param.Email}...))
	return
}

func (u User) LoginHandler(c *gin.Context) {
	// 参数校验
	param := service.UserLoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 业务处理——用户登陆
	svc := service.New(c.Request.Context())
	user, err := svc.Login(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.Login errs: %v", err)
		if err == errcode.ErrorUserNotExit {
			response.ToErrorResponse(errcode.ErrorUserNotExit)
			return
		}
		if err == errcode.ErrorUserPassword {
			response.ToErrorResponse(errcode.ErrorUserPassword)
			return
		}
		response.ToErrorResponse(errcode.ErrorUserLoginFail.WithDetails(err.Error()))
		return
	}

	// 业务响应
	response.ToErrorResponse(errcode.Success.
		WithDetails("username: " + user.UserName).
		WithDetails("accessToken: " + user.AccessToken).
		WithDetails("refreshToken: " + user.RefreshToken))
	return
}

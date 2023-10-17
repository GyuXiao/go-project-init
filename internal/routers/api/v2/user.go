package v2

import (
	"GyuBlog/global"
	"GyuBlog/internal/service"
	"GyuBlog/pkg/app"
	"GyuBlog/pkg/errcode"
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

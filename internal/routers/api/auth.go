package api

import (
	"GyuBlog/global"
	"GyuBlog/internal/service"
	"GyuBlog/pkg/app"
	"GyuBlog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	// 常规的参数校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}
	svc := service.New(c.Request.Context())
	// Auth 鉴权
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	// 生成 token
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}

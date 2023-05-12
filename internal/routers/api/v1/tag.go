package v1

import (
	"GyuBlog/global"
	"GyuBlog/internal/service"
	"GyuBlog/pkg/app"
	"GyuBlog/pkg/convert"
	"GyuBlog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

func (t Tag) List(c *gin.Context) {
	// 验证 validator 是否正常（入参校验和绑定）
	//param := service.TagListRequest{}
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof= 0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 获取标签总数
	//svc := service.New(c.Request.Context())
	//totalTagCnt, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	//if err != nil {
	//	global.Logger.Errorf("svc.CountTag err: %v", err)
	//	response.ToErrorResponse(errcode.ErrorCountTagFail)
	//	return
	//}
	//// 拿到标签列表
	//pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	//tags, err := svc.GetTagList(&param, &pager)
	//if err != nil {
	//	global.Logger.Errorf("svc.GetTagList err: %v", err)
	//	response.ToErrorResponse(errcode.ErrorGetTagListFail)
	//	return
	//}
	// 序列化结果
	//response.ToResponseList(tags, totalTagCnt)
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Create(c *gin.Context) {
	// 验证 validator 是否正常（入参校验和绑定）
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 创建标签
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Update(c *gin.Context) {
	// 验证 validator 是否正常（入参校验和绑定）
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 更新标签
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Delete(c *gin.Context) {
	// 验证 validator 是否正常（入参校验和绑定）
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 删除标签
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

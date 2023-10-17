package api

import (
	"GyuBlog/global"
	"GyuBlog/internal/service"
	"GyuBlog/pkg/app"
	"GyuBlog/pkg/convert"
	"GyuBlog/pkg/errcode"
	"GyuBlog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

/*
调试:
C:\Users\ZK.xiao>curl -X POST http://127.0.0.1:8000/upload/file -F file=@D:/GolandProjects/GyuBlog/logo.jpg -F type=1
{"file_success_url":"http://127.0.0.1:8000/static/96d6f2e7e1f705ab5e59c84a6dc009b2.jpg"}
*/

// UploadFileHandler 上传文件
func (u Upload) UploadFileHandler(c *gin.Context) {
	response := app.NewResponse(c)
	// 读取入参 file 字段的上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	// 利用入参 type 字段作为所上传文件类型的依据
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 通过入参检查后进行 service 的调用，完成文件上传和保存，返回文件的展示地址
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_success_url": fileInfo.AccessUrl,
	})
}

package routers

import (
	_ "GyuBlog/docs"
	"GyuBlog/global"
	"GyuBlog/internal/routers/api"
	v1 "GyuBlog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 测试一下 swagger
	// 可以尝试访问一下：http://127.0.0.1:8000/swagger/index.html
	// url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	tag := v1.NewTag()
	article := v1.NewArticle()

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)       // 新增标签
		apiv1.DELETE("/tags/:id", tag.Delete) // 删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)    // 更新指定标签
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List) // 获取标签列表

		apiv1.POST("/articles", article.Create)       // 新增文章
		apiv1.DELETE("/articles/:id", article.Delete) // 删除指定文章
		apiv1.PUT("/articles/:id", article.Update)    // 更新指定文章
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get) // 获取指定文章
		apiv1.GET("/articles", article.List)    // 获取文章列表
	}

	return r
}

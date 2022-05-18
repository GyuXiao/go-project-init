package routers

import (
	v1 "GyuBlog/internal/routers/api/v1"
	_ "github.com/GyuXiao/GyuBlog/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()

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

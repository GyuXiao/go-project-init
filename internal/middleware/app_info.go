package middleware

import "github.com/gin-gonic/gin"

// 在进程内上下文里设置一些内部信息

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog_service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}

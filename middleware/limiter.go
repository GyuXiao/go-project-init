package middleware

import (
	"GyuBlog/pkg/app"
	"GyuBlog/pkg/errcode"
	"GyuBlog/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIf) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// TakeAvailable 方法会占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，
			// 如果没有可用的令牌，将会返回 0，也就是已经超出配额了，因此这时候我们将返回 errcode.TooManyRequest 状态，告诉客户端需要减缓并控制请求速度
			cnt := bucket.TakeAvailable(1)
			if cnt == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

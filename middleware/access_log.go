package middleware

import (
	"GyuBlog/global"
	"GyuBlog/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// 访问日志记录

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化 AccessLogWriter，将其赋予给当前的 Writer 写入流
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		// request:  当前的请求参数
		// response: 当前的请求结果响应主体
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		// method: 当前的调用方法
		// status: 当前的响应结果状态码
		// begin_time: 调用方法的开始时间
		// end_time: 调用方法的结束时间
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}

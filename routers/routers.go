package routers

import (
	_ "GyuBlog/docs"
	"GyuBlog/global"
	v2 "GyuBlog/handlers/user/v2"
	"GyuBlog/middleware"
	"GyuBlog/pkg/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimitBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))

	user := v2.NewUser()

	//r.GET("/auth", api.GetAuth)
	apiv2 := r.Group("/api/v2")

	// 用户模块
	// 用户注册
	apiv2.POST("/signup", user.SignupHandler)
	// 用户登陆
	apiv2.POST("/login", user.LoginHandler)

	//apiv2.Use(app.JWT())
	return r
}

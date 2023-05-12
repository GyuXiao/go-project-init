package main

import (
	"GyuBlog/global"
	"GyuBlog/internal/middleware"
	"GyuBlog/internal/model"
	"GyuBlog/internal/routers"
	"GyuBlog/pkg/logger"
	"GyuBlog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// init 方法的主要作用是进行应用程序的初始化流程控制，整个应用代码里只有一个 init 方法，在这里调用初始化配置的方法，使配置文件内容映射到应用配置结构体
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	setupValidator()
}

// 在 setupLogger 函数内部对 global 的包全局变量 Logger 进行了初始化，
// 使用 lumberjack 作为日志库的 io.Writer，并且设置日志文件所允许的最大占用空间为 600 MB，日志文件最大生存周期为 10 天，并且设置日志文件名的时间格式为本地时间

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	// ReadSection(k, v)，k 和 v 是一一对应的
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupValidator() {
	uni := ut.New(zh.New())
	middleware.Trans, _ = uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*val.Validate)
	if ok {
		_ = zhTranslations.RegisterDefaultTranslations(v, middleware.Trans)
	}
}

// @title Gyu 博客系统
// @version 1.0
// @description 使用 Go 搭建一个 Blog
// @termsOfService https://github.com/GyuXiao/GyuBlog
func main() {
	// 把映射好的配置和 gin 的运行模式进行配置
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// 校验配置是否真正的映射到配置结构体
	//log.Println(global.ServerSetting)
	//log.Println(global.AppSetting)
	//log.Println(global.DatabaseSetting)

	// 测试 Logger 是否达到预期
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")

	s.ListenAndServe()
}

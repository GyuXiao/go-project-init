package main

import (
	"GyuBlog/global"
	"GyuBlog/internal/model"
	"GyuBlog/internal/routers"
	"GyuBlog/pkg/logger"
	"GyuBlog/pkg/setting"
	"github.com/gin-gonic/gin"
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

// @title Gyu 博客系统
// @version 1.0
// @description 使用 Go 搭建一个 Blog
// @termsOfService https://github.com/GyuXiao/GyuBlog
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//log.Println(global.ServerSetting)
	//log.Println(global.AppSetting)
	//log.Println(global.DatabaseSetting)

	// 测试 Logger 是否达到预期
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")

	s.ListenAndServe()
}

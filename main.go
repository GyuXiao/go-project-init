package main

import (
	"GyuBlog/global"
	"GyuBlog/model"
	"GyuBlog/pkg/logger"
	"GyuBlog/pkg/setting"
	"GyuBlog/pkg/snowflake"
	routers2 "GyuBlog/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	setupValidator()

	// 雪花算法生成分布式 ID
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
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

	// setting 的工作，就是读取配置文件 conf/config.yaml，然后将配置文件的每个模块的内容 unmarshal 给对应的结构体
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
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	model.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupValidator() {
	uni := ut.New(en.New())
	global.Trans, _ = uni.GetTranslator("en")
	v, ok := binding.Validator.Engine().(*val.Validate)
	if ok {
		_ = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
	}
}

func main() {
	// 把映射好的配置和 gin 的运行模式进行配置
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers2.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	defer model.CloseDBEngine()

	// 测试 Logger 是否达到预期
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")

	s.ListenAndServe()
}

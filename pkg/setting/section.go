// Package setting
/**
  @author: zk.xiao
  @date: 2022/5/16
  @note: 声明配置属性的结构体并编写读取区段配置的配置方法
**/
package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

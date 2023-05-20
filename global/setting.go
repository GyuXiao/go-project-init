// Package global
/**
  @author: zk.xiao
  @date: 2022/5/16
  @note: 在读取文件的配置信息后，还需要将配置信息和应用程序关联起来，才能够去使用它
**/
package global

import (
	"GyuBlog/pkg/logger"
	"GyuBlog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)

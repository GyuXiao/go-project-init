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

// 对最初预估的三个区段配置，进行全局变量的声明
// 便于在接下来的步骤将其关联起来，并且提供给应用程序内部调用
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)

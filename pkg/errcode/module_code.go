// Package errcode
/**
  @author: zk.xiao
  @date: 2022/6/1
  @note:
**/
package errcode

var (
	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorUserSignupFail = NewError(20040001, "用户注册失败")
	ErrorUserExit       = NewError(20040002, "用户已经存在")
)

// Package errcode
/**
  @author: zk.xiao
  @date: 2022/6/1
  @note:
**/
package errcode

var (
	ErrorGetTagListFail = NewError(200100001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(200100002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(200100003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(200100004, "删除标签失败")
	ErrorCountTagFail   = NewError(200100005, "统计标签失败")
	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)

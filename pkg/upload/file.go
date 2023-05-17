package upload

import (
	"GyuBlog/global"
	"GyuBlog/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

// 获取文件名称，先是通过获取文件后缀并筛出原始文件名进行 MD5 加密，最后返回经过加密处理后的文件名

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMd5(fileName)
	return fileName + ext
}

// 获取文件后缀

func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取文件保存地址

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 检查保存目录是否存在

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查文件后缀是否包含在约定的后缀配置项中

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			allowExt = strings.ToUpper(allowExt)
			if allowExt == ext {
				return true
			}
		}
	}
	return false
}

// 检查文件大小是否超出最大大小限制

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查文件权限是否足够，和 CheckSavePath 的逻辑相似

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建在文件上传时所需要的保存目录

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// 保存所上传的文件

func SaveFile(file *multipart.FileHeader, dst string) error {
	// 打开源文件的地址
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标地址的文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

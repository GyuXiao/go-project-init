package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 可以将上传的文件名格式化

func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
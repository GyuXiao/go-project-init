package util

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "Gyu.vip"

func EncodeMd5(data []byte) string {
	m := md5.New()
	m.Write([]byte(secret))
	return hex.EncodeToString(m.Sum(data))
}

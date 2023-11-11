package util

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// bcrypt 加密

func EncodeBcrypt(pwdStr string) (string, error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// bcrypt 解密

func DecodeBcrypt(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}

const secret = "Gyu.vip"

// md5 加密

func EncodeMd5(data []byte) string {
	m := md5.New()
	m.Write([]byte(secret))
	return hex.EncodeToString(m.Sum(data))
}

// md5 解密

func DecodeMd5(encryptedString string) string {
	decrypted, _ := hex.DecodeString(encryptedString)
	return string(decrypted)
}

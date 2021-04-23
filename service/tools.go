package service

import (
	"crypto/md5"
	"encoding/hex"
)

//md5String 计算字符串的MD5
func md5String(plainText string) string {
	data := []byte(plainText)
	binaryData := md5.Sum(data)
	return hex.EncodeToString(binaryData[:])
}

//encodePassword 密码加密算法
func encodePassword(username, password string) string {
	return "0x" + md5String(username+password)
}

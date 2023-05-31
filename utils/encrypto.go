package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//EncodePassword 用户名和密码 MD5算法
func EncodePassword(username,password string) string {

	hash := md5.New()
	hash.Write([]byte(username))
	sum := hash.Sum([]byte(password))
	return hex.EncodeToString(sum)

}


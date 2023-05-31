package common

import (
	"fmt"
	"hechuangfil/conf"
	"time"
)

func GenAuthCodeKey(phone string) string {
	return fmt.Sprintf("authCode:%s", phone)
}

func GenTokenKey(token string) string {
	return conf.TOKEN_PREFIX_USER + token
}

func TimeToDay(time time.Time) int32 {
	year := time.Year()
	month := time.Month()
	day := time.Day()
	return int32(year) * 10000 + int32(month) * 100 + int32(day)
}
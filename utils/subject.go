package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"hechuangfil/conf"
	"hechuangfil/entity"
	"log"
	"strconv"
)

//GetSubject 获取当前用户信息,直接从缓存中获取
func GetSubject(client *redis.Client, ctx *gin.Context) (*entity.Customer, error) {
	var tokenStr string
	tokenStr = ctx.GetHeader(conf.API_KEY)
	if tokenStr ==""{
		tokenStr = ctx.DefaultQuery("token", "")
	}
	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(conf.TOKEN_PREFIX_USER+tokenStr).Bytes()
	if err != nil {
		return nil, err
	}
	var sysUser = new(entity.Customer)
	err = json.Unmarshal(userInfoBytes, sysUser)
	return sysUser, err
}

//直接用tokeen串从缓存拿
func GetSubjectByTokenStr(tokenStr string, client *redis.Client,tokentype string) (*entity.Customer, error) {

	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(tokentype + tokenStr).Bytes()
	if err != nil {
		return nil, err
	}
	var sysUser = new(entity.Customer)
	err = json.Unmarshal(userInfoBytes, sysUser)
	return sysUser, err

}

// 获取code,直接从缓存中获取
//func GetCodePhone(client *redis.Client,code string,toentype string) (string) {
//	//从缓存中取到当前登录人信息
//	phone := client
//
//	return sysUser, err
//}

func GetUserSubject(client *redis.Client, ctx *gin.Context) (entity.Customer, error) {
	var tokenStr string
	tokenStr = ctx.GetHeader(conf.API_KEY)
	if tokenStr ==""{
		tokenStr = ctx.DefaultQuery("token", "")
	}
	var User entity.Customer
	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(conf.TOKEN_PREFIX_USER + tokenStr).Bytes()
	if err != nil {
		return User, err
	}
	//var User = new(entity.Customer)

	err = json.Unmarshal(userInfoBytes, &User)
	return User, err
}

func GetUserId(client *redis.Client, ctx *gin.Context) (string) {
	var tokenStr string
	tokenStr = ctx.GetHeader(conf.API_KEY)
	if tokenStr ==""{
		tokenStr = ctx.DefaultQuery("token", "")
	}
	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(conf.TOKEN_PREFIX_USER + tokenStr).Bytes()
	if err != nil {
		log.Println(err)
		return ""
	}
	var User = new(entity.Customer)
	err = json.Unmarshal(userInfoBytes, User)
	if err!=nil{
		log.Println(err)
		return ""
	}

	return strconv.FormatInt(User.CustomerID,10)
}
func GetReferrerId(client *redis.Client, ctx *gin.Context) (string) {
	var tokenStr string
	tokenStr = ctx.GetHeader(conf.API_KEY)
	if tokenStr ==""{
		tokenStr = ctx.DefaultQuery("token", "")
	}
	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(conf.TOKEN_PREFIX_USER + tokenStr).Bytes()
	if err != nil {
		log.Println(err)
		return ""
	}
	var User = new(entity.Customer)
	err = json.Unmarshal(userInfoBytes, User)
	if err!=nil{
		log.Println(err)
		return ""
	}

	return strconv.FormatInt(User.ReferrerId,10)
}

func GetWithdrawWallet(client *redis.Client, ctx *gin.Context) (string) {
	var tokenStr string
	tokenStr = ctx.GetHeader(conf.API_KEY)
	if tokenStr ==""{
		tokenStr = ctx.DefaultQuery("token", "")
	}
	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(conf.TOKEN_PREFIX_USER + tokenStr).Bytes()
	if err != nil {
		log.Println(err)
		return ""
	}
	var User = new(entity.Customer)
	err = json.Unmarshal(userInfoBytes, User)
	if err!=nil{
		log.Println(err)
		return ""
	}

	return User.WithdrawWallet
}

//设置充币钱包
func SetWithdrawWallet(client *redis.Client, ctx *gin.Context,wallet string) error {
	var tokenStr string
	tokenStr = ctx.GetHeader(conf.API_KEY)
	if tokenStr ==""{
		tokenStr = ctx.DefaultQuery("token", "")
	}
	//从缓存中取到当前登录人信息
	userInfoBytes, err := client.Get(conf.TOKEN_PREFIX_USER + tokenStr).Bytes()
	if err != nil {
		return err
	}
	var User = new(entity.Customer)
	err = json.Unmarshal(userInfoBytes, User)
	if err!=nil{
		return err
	}
	User.WithdrawWallet = wallet
	info,err := json.Marshal(User)

	return client.Set(conf.TOKEN_PREFIX_USER+ tokenStr,string(info),conf.TOKEN_EFFECT_TIME).Err()
}

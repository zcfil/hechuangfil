package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hechuangfil/conf"
	"hechuangfil/models"
	"hechuangfil/result"
	"hechuangfil/utils"
	"log"
	"strconv"
	"time"
)


type User struct {
	*models.UserModels
}

func NewUser(u models.UserModels) (*User){
	return &User{
		&u,
	}
}

//@description admin退出
//@accept json
//@Success 200 {object} gin.H
//@router /user/logout [post]
//@Security ApiKeyAuth
func (user *User) DoLogout(ctx *gin.Context) {

	tokenStr := ctx.GetHeader(conf.API_KEY)
	//检查是否携带token
	if tokenStr == "" {

		ctx.String(403, "token为空!")
		ctx.Abort()
		return
	}
	//查询token是否存在缓存里
	if exist, err := user.Exists(conf.TOKEN_PREFIX_ADMIN + tokenStr).Result(); exist == 0 || err != nil {

		log.Println("exist:", exist)
		ctx.String(403, "token已过期!")
		ctx.Abort()
		return
	}

	if result, err := user.Del(conf.TOKEN_PREFIX_ADMIN + tokenStr).Result(); result == 0 || err != nil {

		log.Println("result:", result)
		ctx.String(401, "退出失败:%v",err)
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "data": nil, "msg": "success"})
}

//用户退出
func (user *User) UserLogout(ctx *gin.Context) {

	tokenStr := ctx.GetHeader(conf.API_KEY)
	//检查是否携带token
	if tokenStr == "" {

		ctx.String(403, "token为空!")
		ctx.Abort()
		return
	}
	//查询token是否存在缓存里
	if exist, err := user.Exists(conf.TOKEN_PREFIX_USER + tokenStr).Result(); exist == 0 || err != nil {

		log.Println("exist:", exist)
		ctx.String(403, "token已过期!")
		ctx.Abort()
		return
	}

	if result, err := user.Del(conf.TOKEN_PREFIX_USER + tokenStr).Result(); result == 0 || err != nil {

		log.Println("result:", result)
		ctx.String(401, "退出失败:%v",err)
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "data": nil, "msg": "success"})
}

//个人信息
func (u *User) MemberInfo(ctx *gin.Context){
	fmt.Println("Path:",ctx.Request.URL.Path)
	user, err := utils.GetUserSubject(u.Client, ctx)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	info := struct {
		CustomerId         int64
		Phone   	string
		CreateTime time.Time
		Wallet		string
		Balance		float64
	}{
		CustomerId:         user.CustomerID,
		Phone:   	user.Phone,
		CreateTime: user.CreateTime,
		Wallet:		user.Wallet,
		Balance: 	user.Balance,
	}

	ctx.JSON(200,result.Ok(info))
}

//修改信息
func (u *User)UpdateMember (ctx *gin.Context){
	param := make(map[string]string)
	param["phone"] = ctx.Request.FormValue("phone")
	param["email"] = ctx.Request.FormValue("email")
	param["certification"] = ctx.Request.FormValue("certification")
	for k,v := range param{
		if v ==""{
			ctx.JSON(200, result.Fail(errors.New(k+"不能为空！")))
			ctx.Abort()
			return
		}
	}
	//param["token"] = conf.TOKEN_PREFIX_USER+ctx.GetHeader(conf.API_KEY)
	user, err := utils.GetUserSubject(u.Client, ctx)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	param["id"] = strconv.FormatInt(user.CustomerID,10)
	err = u.UpdateUserInfo(param)
	if err!=nil{
		ctx.JSON(200,result.Fail(err))
	}
	user.Phone = param["phone"]
	//写回Redis
	info,err := json.Marshal(user)

	u.Client.Set(conf.TOKEN_PREFIX_USER+ctx.GetHeader(conf.API_KEY),string(info),conf.TOKEN_EFFECT_TIME)

	ctx.JSON(200,gin.H{"code": 0, "data": err, "msg": "success"})
}


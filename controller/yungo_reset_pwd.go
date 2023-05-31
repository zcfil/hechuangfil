package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hechuangfil/result"
	"hechuangfil/utils"
)

type Restpwd struct{
	OldPwd string `json:"oldPwd" binding:"required"`
	NewPwd string	`json:"newPwd" binding:"required"`
}

//修改密码
func (y *Yungo) ReSetPwd(ctx *gin.Context){

	var params Restpwd
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	user, _ := utils.GetSubject(y.Client, ctx)
	//原密码校验
	oldpwd := utils.EncodePassword(user.Phone, params.OldPwd)
	if user.Password != oldpwd {
		ctx.JSON(200, result.Fail(errors.New("密码错误")))
		ctx.Abort()
		return
	}
	//新密码替换
	newpwd := utils.EncodePassword(user.Phone, params.NewPwd)
	user.Password=newpwd
	_, err = y.Engine.Id(user.CustomerID).Cols("password").Update(user)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	//修改成功
	ctx.JSON(200,result.Ok(nil))
}



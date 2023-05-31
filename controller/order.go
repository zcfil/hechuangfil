package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hechuangfil/result"
	"hechuangfil/utils"
	"strconv"
)

//获取订单列表
func (u *User)Orderlist(ctx *gin.Context){
	user, err := utils.GetUserSubject(u.Client, ctx)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	param := make(map[string]string)
	param["pageNo"] = ctx.DefaultQuery("pageIndex","1")
	param["pageSize"] = ctx.DefaultQuery("pageSize","10")
	//param["status"] = ctx.Request.FormValue("status")
	param["user_id"] = strconv.FormatInt(user.CustomerID,10)

	res ,total, err := u.GetOrderList(param)
	if err!=nil{
		ctx.JSON(200,result.Fail(err))
		return
	}
	ctx.JSON(200,utils.NewPageDataTotal(param,res,total))
}

//下单
func (u *User)PlaceAnOrder(ctx *gin.Context){
	param := make(map[string]string)

	param["hashrate"] = ctx.Request.FormValue("hashrate")
	param["remark"] = ctx.Request.FormValue("remark")
	param["userid"] = utils.GetUserId(u.Client,ctx)
	param["order_id"] = strconv.FormatInt(int64(uuid.New().ID()),10)
	amount ,err := u.GetHashrateAmount(param["hashrate"])
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	if amount > u.GetBalance(param["userid"] ){
		ctx.JSON(200, result.Fail(errors.New("余额不足")))
		ctx.Abort()
		return
	}
	price,_:= u.GetHashratePrice()
	param["price"]  = utils.Float64ToString(price)
	param["amount"] = utils.Float64ToString(amount)
	param["referrer_id"] = utils.GetReferrerId(u.Client,ctx)
	err = u.PlaceTheOrder(param)
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	res := make(map[string]string)
	res["order_id"] = param["order_id"]

	ctx.JSON(200,result.Ok(res))
}

//下单
func (u *User)TashratePrice(ctx *gin.Context){

	price,err := u.GetHashratePrice()
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	ctx.JSON(200,result.Ok(price))
}

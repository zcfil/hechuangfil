package controller

import (
	"github.com/gin-gonic/gin"
	"hechuangfil/result"
	"hechuangfil/utils"
	"net/http"
)
//
func (u *User)GetCollectInfo2(ctx *gin.Context){
	userid := utils.GetUserId(u.Client,ctx)
	res ,err := u.CollectInfo2(userid)
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	ctx.JSON(200,result.Ok(res))
}
func (u *User)OrderProfitList(c *gin.Context) {
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["user_id"] = utils.GetUserId(u.Client, c)

	dataList , err := u.GetOrderProfitList(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}
	res := utils.NewPageData(param,dataList)

	c.JSON(http.StatusOK,res)
}


func (u *User)OrderProfitById(c *gin.Context) {
	param := make(map[string]string)

	param["user_id"] = utils.GetUserId(u.Client, c)
	param["id"] = c.Request.FormValue("id")

	var data map[string]string
	var err error
	//分润收益
	data , err = u.GetOrderProfitById(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,result.Ok(data))
}

func (this *User) LowerLevelPerformance(c *gin.Context) {
	userid := utils.GetUserId(this.Client,c)
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空"))
		c.Abort()
		return
	}

	param := make(map[string]string)
	param["pageSize"] = pageSize
	param["pageIndex"] = pageIndex

	data, err := this.LowerPowerPerformance(userid, param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	pageData := utils.NewPageData(param, data)
	c.JSON(http.StatusOK, pageData)
}
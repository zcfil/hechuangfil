package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hechuangfil/common"
	"hechuangfil/conf"
	"hechuangfil/entity"
	log "hechuangfil/logrus"
	"hechuangfil/result"
	"hechuangfil/utils"
	"net/http"
	"strconv"
)

//充值
func (u *User)RechargeCoin(ctx *gin.Context){
	param := make(map[string]string)

	param["amount"] = ctx.Request.FormValue("amount")
	param["userid"] = utils.GetUserId(u.Client,ctx)
	err := u.RechargeCoinFil(param)
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	res := make(map[string]string)
	res["orderid"] = param["orderid"]

	ctx.JSON(200,result.Ok(res))
}

//提现
func (u *User)WithdrawCoin(ctx *gin.Context){
	param := make(map[string]string)
	password := ctx.Request.FormValue("password")
	if password ==""{
		ctx.JSON(200, result.Failstr("密码为空!"))
		ctx.Abort()
		return
	}
	user,_ := utils.GetUserSubject(u.Client,ctx)
	param["to_address"] = user.WithdrawWallet

	if user.PayPassword != utils.EncodePassword(user.Phone, password){
		ctx.JSON(200, result.Failstr("密码错误!"))
		ctx.Abort()
		return
	}

	param["userid"] = utils.GetUserId(u.Client,ctx)

	param["amount"] = ctx.Request.FormValue("amount")
	amount,err := strconv.ParseFloat(param["amount"],64)
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	if u.GetBalance(param["userid"]) < amount{
		ctx.JSON(200, result.Failstr("余额不足"))
		ctx.Abort()
		return
	}

	if err := u.WithdrawCoinFil(param);err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}


	ctx.JSON(200,result.Ok("提交成功！"))
}

// 充值列表
func (u *User)GetRechargeList(c *gin.Context) {
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["user_id"] = utils.GetUserId(u.Client, c)

	dataList , err := u.RechargeList(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}
	res := utils.NewPageData(param,dataList)

	c.JSON(http.StatusOK,res)
}
//充值记录
func (u *User)GetRechargeById(c *gin.Context) {
	param := make(map[string]string)

	param["user_id"] = utils.GetUserId(u.Client, c)
	param["recharge_id"] = c.Request.FormValue("recharge_id")
	data , err := u.RechargeById(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,result.Ok(data))
}

// 提现列表
func (u *User) GetWithdrawList(c *gin.Context) {
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["user_id"] = utils.GetUserId(u.Client, c)

	dataList , err := u.WithdrawList(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}
	res := utils.NewPageData(param,dataList)

	c.JSON(http.StatusOK, res)
}
// 提现记录
func (u *User) GetWithdrawById(c *gin.Context) {
	param := make(map[string]string)
	param["withdraw_id"] = c.Request.FormValue("withdraw_id")
	param["user_id"] = utils.GetUserId(u.Client, c)

	data , err := u.WithdrawById(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Ok(data))
}
//充币钱包
func (u *User)GetWalletAddress(ctx *gin.Context){
	userid := utils.GetUserId(u.Client,ctx)
	res ,err := u.WalletBalance(userid)
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}


	ctx.JSON(200,result.Ok(res))
}

//总有效算力，锁仓总额，质押总额，今日收益，总收益
func (u *User)GetCollectInfo(ctx *gin.Context){
	userid := utils.GetUserId(u.Client,ctx)
	res ,err := u.CollectInfo(userid)
	if err!=nil{
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	ctx.JSON(200,result.Ok(res))
}

//获取提现钱包
func (u *User)GetWithdrawAddress(ctx *gin.Context){
	ctx.JSON(200,result.Ok(utils.GetWithdrawWallet(u.Client,ctx)))
}

//修改提现钱包
func (u *User)UpdateWithdrawAddress(c *gin.Context){
	param := make(map[string]string)
	param["wallet"] = c.Request.FormValue("wallet")
	if param["wallet"]==""{
		c.JSON(http.StatusOK,result.Failstr("钱包地址不能为空！"))
		c.Abort()
		return
	}
	param["user_id"] = utils.GetUserId(u.Client, c)
	var err error
	sess := u.Engine.NewSession()
	sess.Begin()
	err = u.SetWithdrawWallet(sess,param)
	if err != nil {
		sess.Rollback()
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	if err = utils.SetWithdrawWallet(u.Client,c,param["wallet"]);err!=nil{
		sess.Rollback()
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	sess.Commit()
	c.JSON(200,result.Ok("修改成功！"))
}

// ChangePayPasswordByCode 修改支付密码
func (u *User) ChangePayPasswordByCode(c *gin.Context) {
	tokenStr := c.GetHeader(conf.API_KEY)
	tokenKey := common.GenTokenKey(tokenStr)

	newPayPassword := c.Request.FormValue("newPayPassword")
	if newPayPassword == "" {
		c.JSON(http.StatusOK, result.Fail(errors.New("新的支付密码不能为空")))
		c.Abort()
		return
	}

	code := c.Request.FormValue("code")
	if code == "" {
		c.JSON(http.StatusOK, result.Fail(errors.New("验证码不能为空")))
		c.Abort()
		return
	}

	infoStr, err := u.Client.Get(tokenKey).Result()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	customerInfo := new(entity.Customer)
	err = json.Unmarshal([]byte(infoStr), customerInfo)
	if err != nil {
		log.Error("解析顾客信息报错，err:",err.Error())
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	codeKey := common.GenAuthCodeKey(customerInfo.Phone)
	authCode, err := u.Client.Get(codeKey).Result()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	if authCode != code {
		c.JSON(http.StatusOK, result.Fail(errors.New("验证码不正确")))
		c.Abort()
		return
	}
	user,_ := utils.GetUserSubject(u.Client,c)
	//phone := user.Phone + newPayPassword
	newPayPassword = utils.EncodePassword(user.Phone, newPayPassword)

	if err := u.changePayPassword(tokenKey, newPayPassword, customerInfo); err != nil {
		c.JSON(http.StatusOK, result.Fail(errors.New("修改支付密码失败")))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))

}

func (u *User) ChangePayPasswordByPassword(c *gin.Context) {
	tokenStr := c.GetHeader(conf.API_KEY)
	tokenKey := common.GenTokenKey(tokenStr)

	newPayPassword := c.Request.FormValue("newPayPassword")
	if newPayPassword == "" {
		c.JSON(http.StatusOK, result.Fail(errors.New("新的支付密码不能为空")))
		c.Abort()
		return
	}

	password := c.Request.FormValue("password")
	if password == "" {
		c.JSON(http.StatusOK, result.Fail(errors.New("登录密码不能为空")))
		c.Abort()
		return
	}

	infoStr, err := u.Client.Get(tokenKey).Result()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	customerInfo := new(entity.Customer)
	err = json.Unmarshal([]byte(infoStr), customerInfo)
	if err != nil {
		log.Error("解析顾客信息报错，err:",err.Error())
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	encodePassword := utils.EncodePassword(customerInfo.Phone, password)
	if encodePassword != customerInfo.Password {
		c.JSON(http.StatusOK, result.Fail(errors.New("密码不正确")))
		c.Abort()
		return
	}
	newPayPassword = utils.EncodePassword(customerInfo.Phone, newPayPassword)
	if err := u.changePayPassword(tokenKey, newPayPassword, customerInfo); err != nil {
		c.JSON(http.StatusOK, result.Fail(errors.New("修改支付密码失败")))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}

func (u *User) changePayPassword(tokenKey, newPayPassword string, customerInfo *entity.Customer) (err error) {
	sql := `update customer set pay_password = '%s' where customer_id = %d`
	sql = fmt.Sprintf(sql, newPayPassword, customerInfo.CustomerID)

	session := u.Engine.NewSession()
	err = session.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			session.Rollback()
		}
		session.Close()
	}()

	 _, err = session.Exec(sql)
	 if err != nil {
		return
	}

	customerInfo.PayPassword = newPayPassword
	dataBin, _ := json.Marshal(customerInfo)
	err = u.Client.Set(tokenKey, dataBin, conf.TOKEN_EFFECT_TIME).Err()
	if err != nil {
		return
	}
	err = session.Commit()
	if err != nil {
		return
	}
	return
}

func (u *User) GetCustomerPhone(c *gin.Context) {
	tokenStr := c.GetHeader(conf.API_KEY)
	tokenKey := common.GenTokenKey(tokenStr)
	infoStr, err := u.Client.Get(tokenKey).Result()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	customerInfo := new(entity.Customer)
	err = json.Unmarshal([]byte(infoStr), customerInfo)
	if err != nil {
		log.Error("解析顾客信息报错，err:",err.Error())
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Ok(customerInfo.Phone))
}
func (u *User) GetCustomerCode(c *gin.Context) {
	tokenStr := c.GetHeader(conf.API_KEY)
	tokenKey := common.GenTokenKey(tokenStr)
	infoStr, err := u.Client.Get(tokenKey).Result()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	customerInfo := new(entity.Customer)
	err = json.Unmarshal([]byte(infoStr), customerInfo)
	if err != nil {
		log.Error("解析顾客信息报错，err:",err.Error())
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	p,_ := u.GetHashratePrice()
	mylevel,_ := u.GetSetUserLevel(0,p)
	if len(mylevel)==0{
		c.JSON(http.StatusOK, result.Failstr("未设置会员等级"))
		c.Abort()
		return
	}
	acc := u.GetAccumulative(strconv.FormatInt(customerInfo.CustomerID,10))
	if mylevel[0].Accumulative>acc{
		c.JSON(http.StatusOK, result.Fail(fmt.Errorf("被推荐人+自身算力未达到%fT算力！目前为：%fT",mylevel[0].Accumulative,acc)))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Ok(map[string]string{"invitation_code":customerInfo.InvitationCode}))
}

func (u *User)SettleLogList(c *gin.Context) {
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["user_id"] = utils.GetUserId(u.Client, c)

	dataList , err := u.GetSettleLogList(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}
	res := utils.NewPageData(param,dataList)

	c.JSON(http.StatusOK,res)
}

func (u *User)SettleLogById(c *gin.Context) {
	param := make(map[string]string)

	param["user_id"] = utils.GetUserId(u.Client, c)
	param["id"] = c.Request.FormValue("id")
	var data map[string]string
	var err error
	data , err = u.GetSettleLogById(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":-1, "data":nil, "msg":err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,result.Ok(data))
}

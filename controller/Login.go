package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"hechuangfil/common"
	"hechuangfil/define"
	"io/ioutil"
	rand2 "math/rand"
	"strconv"
	"strings"

	"hechuangfil/conf"
	"hechuangfil/entity"
	"hechuangfil/models"
	"hechuangfil/result"
	"hechuangfil/utils"
	"net/http"
	"time"

	log "hechuangfil/logrus"
)

var (
	ErrUnregistered error = errors.New("未注册的手机号")			// 未注册的手机号
)

type Login struct {
	*models.LoginModels
}
func NewLogin(login models.LoginModels) (*Login){
	return &Login{
		&login,
	}
}

//@description 用户登录
//@accept json
//@Param loginDto body dto.LoginDto true "loginDto"
//@Success 200 {object} gin.H
//@router /login [post]
func (login *Login) DoLogin(ctx *gin.Context) {

	//var loginDto dto.LoginDto

	//if err := ctx.ShouldBindJSON(&loginDto); err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": err.Error()})
	//	ctx.Abort()
	//	return
	//
	//}

	username := ctx.Request.FormValue("username")
	password := ctx.Request.FormValue("password")
	sysUser := new(entity.SysUser)
	if ib, err := login.Engine.Where("sys_user.username=? and sys_user.delflag=0", username).Get(sysUser); err != nil || ib == false {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "用户名错误"})
		ctx.Abort()
		return
	}
	//密码b不匹配
	if utils.EncodePassword(username, password) != sysUser.Password {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "密码错误"})
		ctx.Abort()
		return
	}

	//签名生成token,设置有效期为一天
	token, err := utils.Sign(username, int64(conf.TOKEN_EFFECT_TIME))
	if err != nil {
		ctx.String(500, err.Error())
		ctx.Abort()
		return
	}
	//用户基本信息缓存到redis,并设置有效期为token有效期
	userInfo, _ := json.Marshal(sysUser)
	err = login.SetNX(conf.TOKEN_PREFIX_ADMIN+token, string(userInfo), conf.TOKEN_EFFECT_TIME).Err()
	if err != nil {
		ctx.String(500, err.Error())
		ctx.Abort()
		return
	}
	//加密生成token，并存到redis中
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": token, "msg": "ok"})
}

//@description 管理员注册
//@accept json
//@Param sysUser body entity.SysUser true "sysUser"
//@Success 200 {object} gin.H
//@router /register [post]
func (login *Login) DoRegister(ctx *gin.Context) {
	var sysUser entity.SysUser
	//if err := ctx.ShouldBindJSON(&sysUser); err != nil {
	//	ctx.JSON(http.StatusOK, result.Fail(err))
	//	ctx.Abort()
	//	return
	//}
	username := ctx.Request.FormValue("username")
	if username==""{
		ctx.JSON(http.StatusOK, result.Fail(errors.New("用户名不能为空！")))
		ctx.Abort()
	}
	password := ctx.Request.FormValue("password")
	if password==""{
		ctx.JSON(http.StatusOK, result.Fail(errors.New("密码不能为空！")))
		ctx.Abort()
	}
	bank := ctx.Request.FormValue("bank")
	if bank==""{
		bank = "普通管理员"
	}
	//用户名和密码md5加密
	sysUser.Id = utils.Node().Generate().Int64()
	sysUser.Password = utils.EncodePassword(username, password)
	sysUser.CreateDate = time.Now()
	sysUser.Rank = bank
	sysUser.Username = username

	sql := ` SELECT count(*) count FROM sys_user WHERE (username='`+username+`' and delflag = 0 ); `
	res,_ := login.Engine.QueryString(sql)

	if res[0]["count"] != "0"{
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "用户已存在"})
		ctx.Abort()
		return
	}

	if _, err := login.Engine.Insert(&sysUser); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result.Ok(nil))
}

func (login *Login) VerificationCode(c *gin.Context) {
	phone := c.Request.FormValue("phone")
	if len(phone)<11{
		c.JSON(http.StatusOK,result.Fail(errors.New("phone 手机号码格式不对！")))
		c.Abort()
		return
	}
	msgType := c.Request.FormValue("msgType")
	if len(msgType) <= 0{
		c.JSON(http.StatusOK,result.Fail(errors.New("msgType不能为空！")))
		c.Abort()
		return
	}

	templateID := int64(0)
	switch msgType {
	case define.CLIENT_CODE_TYPE_REGISTER:			 // 注册验证码
		if login.isRegister(phone) {
			c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "该手机号已注册"})
			c.Abort()
			return
		}
		templateID = define.TEMPLATE_REGISTER_CODE
	case define.CLIENT_CODE_TYPE_LOGIN:
		if !login.isRegister(phone) {
			c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "该手机号未注册"})
			c.Abort()
			return
		}
		templateID = define.TEMPLATE_LOGIN_CODE
	case define.CLIEN_CODE_TYPE_AUTH:
		if !login.isRegister(phone) {
			c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "该手机号未注册"})
			c.Abort()
			return
		}
		templateID = define.TEMPLATE_AUTH_CODE
	default:
		c.JSON(http.StatusOK, result.Fail(errors.New("msgType错误")))
		c.Abort()
		return
	}

	// 判断一下验证码的时间
	lastExpireTime, err := login.getAuthCodeExpireTime(phone)
	if err != nil {
		if err != redis.Nil {
			c.JSON(http.StatusOK,result.Fail(errors.New("redis错误！")))
			c.Abort()
			log.Error("redis 错误 err:", err.Error())
			return
		}
	}

	if define.AUTH_CODE_VALID_TIME * time.Second - lastExpireTime < time.Second * 60 {
		c.JSON(http.StatusOK,result.Fail(errors.New("请一分钟以后再次发送验证码")))
		c.Abort()
		return
	}

	code := login.randCode(define.AUTH_CODE_WIDTH)
	key := common.GenAuthCodeKey(phone)
	err = login.Client.Set(key, code, time.Second * define.AUTH_CODE_VALID_TIME).Err()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	err = login.requestYouPai(phone, code, templateID)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}


func (login *Login) requestYouPai(phone, code string, templateID int64) error {
	type requestData struct {
		Mobile   	string 		`json:"mobile"`					// 手机号码
		TemplateID  int64 		`json:"template_id"`			// 模版ID
		Variables 	[]string 	`json:"variables"`				// 参数列表
	}

	reqData := &requestData{
		Mobile:     phone,
		TemplateID: templateID,
		Variables:  []string{code},
	}

	reqJson, err := json.Marshal(reqData)
	if err != nil {
		log.Error("整理请求数据错误, err:", err.Error())
		return err
	}
	jsonInfo := strings.NewReader(string(reqJson))

	url := define.AUTH_CODE_URL
	req, err := http.NewRequest("POST", url, jsonInfo)
	if err != nil {
		log.Error("创建请求错误，err:", err.Error())
		return err
	}

	req.Header.Add("Authorization", login.AuthCode.Token)
	req.Header.Add("Content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("访问验证码请求错误, err:", err.Error())
		return err
	}

	defer res.Body.Close()

	// 取出返回值做个验证
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("获取返回数据错误, err:", err.Error())
		return err
	}

	if string(body) == "Unauthorized" {
		return errors.New("token 验证已过期")
	}

	log.Info("body:", string(body))
	mapData := make(map[string]interface{})
	if err := json.Unmarshal(body, &mapData); err != nil {
		log.Error("解析返回数据错误， err:", err)
		return err
	}


	retCode, ok := mapData["code"]
	if !ok {
		return errors.New("返回码错误")
	}

	if retCode.(float64) != 0 {
		return errors.New("返回码错误")
	}
	return nil
}

func (Login *Login) randCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand2.Intn(r) ])
	}
	return sb.String()
}

//@description 商户注册申请
//@accept json
//@Param MerchantsDto body dto.MerchantsDto true "MerchantsDto"
//@Success 200 {object} gin.H
//@router /merchants/register [post]
func (login *Login) CustomerRegister(ctx *gin.Context) {
	var Customer entity.Customer
	phone := ctx.Request.FormValue("phone")
	if phone==""{
		ctx.JSON(http.StatusOK,result.Fail(errors.New("手机号码不能为空!")))
		ctx.Abort()
		return
	}
	code := ctx.Request.FormValue("code")
	if code==""{
		ctx.JSON(http.StatusOK,result.Fail(errors.New("验证码不能为空!")))
		ctx.Abort()
		return
	}
	invitation_code := ctx.Request.FormValue("invitation_code")
	if invitation_code==""{
		ctx.JSON(http.StatusOK,result.Fail(errors.New("验证码不能为空!")))
		ctx.Abort()
		return
	}

	// 验证码暂时没接 先不处理
	authCode, err := login.getAuthCode(phone)
	if err != nil || authCode != code {
		ctx.JSON(http.StatusOK, result.Fail(errors.New("验证码错误或已失效")))
		ctx.Abort()
		return
	}

	if !login.CanRegisterByPhone(phone) {
		ctx.JSON(http.StatusOK,result.Fail(errors.New("手机号码已注册!")))
		ctx.Abort()
		return
	}
	password := ctx.Request.FormValue("password")
	if password==""{
		ctx.JSON(http.StatusOK,result.Fail(errors.New("密码不能为空!")))
		ctx.Abort()
		return
	}

	//用户名和密码md5加密
	Customer.CustomerID = utils.Node().Generate().Int64()
	Customer.Password = utils.EncodePassword(phone, password)
	Customer.CreateTime = time.Now()
	Customer.Phone = phone

	Customer.PayPassword = Customer.Password 		// 支付密码默认设置成登录密码

	session := login.Engine.NewSession()
	var data []map[string]interface{}
	defer func() {
		if err!=nil{
			session.Rollback()
			return
		}
		session.Commit()
	}()
	session.Begin()
	sql := `select customer_id from customer where phone = "` + phone + `" limit 1`
	data, err = session.QueryInterface(sql)
	if err != nil {
		log.Error("Register err:", err.Error())
		ctx.JSON(http.StatusOK,result.Fail(errors.New("注册失败1!"+err.Error())))
		ctx.Abort()
		return
	}

	if len(data) > 0 {
		ctx.JSON(http.StatusOK,result.Fail(errors.New("注册失败2!"+err.Error())))
		ctx.Abort()
		return
	}
	if Customer.Wallet,err = login.Response.WalletNew();err!=nil{
		ctx.JSON(http.StatusOK,result.Fail(errors.New("注册失败3!"+err.Error())))
		ctx.Abort()
		return
	}
	reid := utils.GetInvitationCode(login.Client)[invitation_code]
	if reid==""&&invitation_code!="AAAAAA"{
		ctx.JSON(http.StatusOK,result.Fail(errors.New("邀请码不存在!")))
		ctx.Abort()
		return
	}
	if reid==""{
		reid = "0"
	}
	Customer.ReferrerId,_ = strconv.ParseInt(reid,10,64)

	cd,err := utils.SetInvitationCode(login.Client,strconv.FormatInt(Customer.CustomerID,10))
	if err != nil {
		ctx.JSON(http.StatusOK,result.Fail(errors.New("注册失败5!"+err.Error())))
		ctx.Abort()
		return
	}
	Customer.InvitationCode = cd

	if _, err = session.Insert(&Customer); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	if err = login.SetReferrer(session,Customer.CustomerID,reid);err!=nil{
		ctx.JSON(http.StatusOK, result.Fail(err))
		ctx.Abort()
		return
	}
	filaddr := make(map[string]string)
	filaddr[Customer.Wallet] = strconv.FormatInt(Customer.CustomerID,10)
	if err = utils.AddFilAddress(login.Client,filaddr);err!=nil{
		ctx.JSON(http.StatusOK,result.Failstr("注册失败4:"+ err.Error()+err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result.Ok(nil))
}

func (login *Login) CustomerLoginByCode(ctx *gin.Context) {
	phone := ctx.Request.FormValue("phone")
	if phone ==""{
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "手机号不能为空！"})
		ctx.Abort()
		return
	}
	code := ctx.Request.FormValue("code")
	if code ==""{
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "验证码不能为空！"})
		ctx.Abort()
		return
	}

	customerInfo, err := login.getCustomerInfo(phone)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": err.Error()})
		ctx.Abort()
		return
	}

	//  验证验证码是否正确
	authCode, err := login.getAuthCode(phone)
	if err != nil || authCode != code {
		ctx.JSON(http.StatusOK, result.Fail(errors.New("验证码错误或已失效")))
		ctx.Abort()
		return
	}


	login.customerLogin(ctx, phone, customerInfo)
}

func (login *Login) CustomerLoginByPassword(ctx *gin.Context) {
	phone := ctx.Request.FormValue("phone")
	if phone ==""{
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "手机号不能为空！"})
		ctx.Abort()
		return
	}

	pwd := ctx.Request.FormValue("password")
	if pwd ==""{
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "密码不能为空！"})
		ctx.Abort()
		return
	}

	customerInfo, err := login.getCustomerInfo(phone)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": err.Error()})
		ctx.Abort()
		return
	}
	// 验证密码是否正确
	encodePwd := utils.EncodePassword(phone, pwd)
	//if encodePwd != string(customerInfo["password"].([]byte)) {
	if encodePwd != customerInfo.Password {
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "密码错误！"})
		ctx.Abort()
		return
	}

	login.customerLogin(ctx, phone, customerInfo)
}


func (login *Login) customerLogin(ctx *gin.Context, phone string, customer entity.Customer) {
	//签名生成token,设置有效期为一天
	token, err := utils.Sign(phone, int64(conf.TOKEN_EFFECT_TIME))
	if err != nil {
		ctx.String(500, err.Error())
		ctx.Abort()
		return
	}
	//用户基本信息缓存到redis,并设置有效期为token有效期的2倍
	userInfo, _ := json.Marshal(customer)
	err = login.SetNX(conf.TOKEN_PREFIX_USER+token, string(userInfo), conf.TOKEN_EFFECT_TIME).Err()
	if err != nil {
		ctx.String(500, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": token, "msg": "ok"})
}

func (login *Login) getCustomerInfo(phone string) (customer entity.Customer, err error) {
	ib, err1 := login.Engine.Where("customer.phone = ? ",phone).Get(&customer)
	if err1 !=nil{
		err = err1
		return
	}
	if !ib{
		err = ErrUnregistered
		return
	}
	return
}

func (login *Login) ChangePassword(c *gin.Context) {
	phone := c.Request.FormValue("phone")
	if phone == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "手机号不能为空！"})
		c.Abort()
		return
	}

	password := c.Request.FormValue("password")
	if phone == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "新密码不能为空！"})
		c.Abort()
		return
	}

	code := c.Request.FormValue("code")
	if phone == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "验证码不能为空！"})
		c.Abort()
		return
	}

	// 判断验证码是否正确
	authCode, err := login.getAuthCode(phone)
	if err != nil || authCode != code {
		c.JSON(http.StatusOK, result.Fail(errors.New("验证码错误或已失效")))
		c.Abort()
		return
	}

	encodePassword := utils.EncodePassword(phone, password)

	log.Info("验证码:", code)
	sql := `update customer set password = "%s" where phone = "%s"`
	sql = fmt.Sprintf(sql, encodePassword, phone)

	session := login.NewSession()
	err = session.Begin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "修改密码失败！"})
		c.Abort()
		return
	}

	defer func() {
		if err != nil {
			if err = session.Rollback(); err != nil {
				log.Error("数据回滚失败:", err.Error())
			}
		}
		session.Close()
	}()
	_, err = session.Exec(sql)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "修改密码失败！"})
		c.Abort()
		return
	}

	token := c.GetHeader(conf.API_KEY)
	err = login.changeTokenPassword(token, encodePassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "redis 更新密码失败！"})
		c.Abort()
		return
	}

	err = session.Commit()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": "提交修改密码数据失败！"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}

func (login *Login) changeTokenPassword(token, newPassword string) error {
	tokenKey := common.GenTokenKey(token)
	data, err := login.Client.Get(tokenKey).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return err
	}
	var customer entity.Customer
	err = json.Unmarshal([]byte(data), &customer)
	if err != nil {
		return err
	}

	customer.Password = newPassword
	dataBin, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return login.Client.Set(tokenKey, dataBin, conf.TOKEN_EFFECT_TIME).Err()
}

func (login *Login) getAuthCode(phone string) (string, error) {
	key := common.GenAuthCodeKey(phone)
	return login.Client.Get(key).Result()
}

func (login *Login) getAuthCodeExpireTime(phone string) (time.Duration, error) {
	key := common.GenAuthCodeKey(phone)
	return login.Client.TTL(key).Result()
}

// isRegister  判断手机号码是否已经注册
func (login *Login) isRegister(phone string) bool {
	sql := `select customer_id from customer where phone = "` + phone + `" and is_del = 0 limit 1`
	res, err := login.Engine.QueryString(sql)
	if err != nil {
		log.Error("查找手机号是否已经注册错误, err", err.Error())
		return false
	}
	if len(res) <= 0 {
		return false
	}
	return true
}
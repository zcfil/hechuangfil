package models

import (
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"go.uber.org/dig"
	"hechuangfil/conf"
	"hechuangfil/utils"
	"strconv"
)

type UserModels struct {
	dig.In
	*xorm.Engine
	*redis.Client
	*conf.Repo
	*SysConfigModels
	*OrderProfitModels
}

//修改用户信息
func (y *UserModels) UpdateUserInfo(param map[string]string)error{
	sql := ` update user set email = :email,phone=:phone,certification=:certification where id=:id `
	sql = utils.SqlReplaceParames(sql,param)
	_,err := y.Engine.Exec(sql)
	if err!=nil{
		return err
	}
	return err
}
//获取余额
func (y *UserModels) GetBalance (customerid string)float64{
	sql := ` select balance from customer where customer_id = ? `
	res,err := y.Engine.QueryString(sql,customerid)
	if err!=nil|| len(res)==0{
		return 0
	}
	f,_ := strconv.ParseFloat(res[0]["balance"],64)
	return f
}
//获取用户业绩
func (y *UserModels) GetAccumulative (customerid string)float64{
	sql := ` select accumulative from customer where customer_id = ? `
	res,err := y.Engine.QueryString(sql,customerid)
	if err!=nil|| len(res)==0{
		return 0
	}
	f,_ := strconv.ParseFloat(res[0]["accumulative"],64)
	return f
}
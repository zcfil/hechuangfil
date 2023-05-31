package models

import (
	"github.com/go-xorm/xorm"
	"log"
	"strconv"
)


type SysConfigModels struct {
	//dig.In
	*xorm.Engine
}

const (
	 HASHRATE_PRICE = "hashrate_price"
)
func NewSysConfigModels(x *xorm.Engine)*SysConfigModels{
	return &SysConfigModels{
		x,
	}
}
func (s *SysConfigModels)GetConfigValue(key string)string{
	sql := `select configValue from sys_config where configKey = ? `
	res,_ := s.Engine.QueryString(sql,key)
	if len(res)==0{
		return ""
	}
	return res[0]["configValue"]
}
//获取对应算力需多少钱
func (s *SysConfigModels)GetHashrateAmount(hashrate string)(float64,error){
	//money,_ := strconv.ParseFloat(amount,64)
	rate,err := strconv.ParseFloat(hashrate,64)
	if err!=nil{
		return 0,err
	}
	price,err := strconv.ParseFloat(s.GetConfigValue(HASHRATE_PRICE),64)
	if err!=nil||price==0{
		log.Println(price,err)
		return 0,err
	}

	return rate * price,err
}
//算力单价
func (s *SysConfigModels)GetHashratePrice()(float64,error){
	price,err := strconv.ParseFloat(s.GetConfigValue(HASHRATE_PRICE),64)
	if err!=nil||price==0{
		log.Println(price,err)
		return 0,err
	}

	return price,err
}
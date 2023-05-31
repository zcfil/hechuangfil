package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"hechuangfil/conf"
	"hechuangfil/entity"
	log "hechuangfil/logrus"
	"hechuangfil/utils"
	"strconv"
	"strings"
)

type OrderProfitModels struct {
	*xorm.Engine
	*UserLevelModels
	*conf.Project
}

func NewOrderProfitModels(x *xorm.Engine, userlevel *UserLevelModels, pro *conf.Project) *OrderProfitModels {
	return &OrderProfitModels{
		x,
		userlevel,
		pro,
	}
}

//func (in *OrderProfitModels)AddOrderProfit(session *xorm.Session, param map[string]string,price float64)(float64,error){
//	//var le entity.UserLevel
//	//var re Referrer
//	//获取订单信息
//	//ins ,err := in.GetOrderProfitById(param["investmentid"])
//	//if err!=nil{
//	//	return err
//	//}
//	Referrers,err := in.GetReferrers(param["userid"])
//	if err!=nil{
//		return 0,err
//	}
//	//获取自己等级
//	mylevel,err := in.GetSetUserByUserid(param["userid"],price)
//	if err!=nil{
//		return 0,err
//	}
//	//获取上级等级
//	Referrers = strings.Trim(Referrers,",")
//	var ulevel [] entity.UserLevel
//
//	if Referrers!=""{
//		ulevel,err = in.GetReferrerLevel(Referrers,price)
//		if err!=nil{
//			return 0,err
//		}
//		//ulevel = append(ulevel, le...)
//	}
//
//	//获取设置等级
//	setlevel,err := in.GetSetUserLevel(mylevel.Levelvalue,price)
//	//升级所需
//	type upgrade struct{
//		need float64
//		LevelvaluePrice float64
//		percent	float64
//	}
//	var up []upgrade
//	//计算自己的利润
//	var isps []entity.OrderProfit
//	var ip entity.OrderProfit
//	orderamount,_ := strconv.ParseFloat(param["amount"],64)
//	camount := orderamount
//	ip.UserId = param["userid"]
//	ip.OrderId = param["order_id"]
//	ip.CustomerId = param["userid"]
//
//	for i:=len(setlevel)-1; i>=0;i--{
//		if mylevel.AccumulativePrice + orderamount > setlevel[i].LevelvaluePrice {
//			//平级
//			if mylevel.AccumulativePrice > setlevel[i].LevelvaluePrice{
//				ip.Profits += camount * setlevel[i].Percentreality
//				break
//			}
//			//越级
//			a :=  mylevel.AccumulativePrice + camount - setlevel[i].LevelvaluePrice
//			ip.Profits += a * setlevel[i].Percentreality
//			camount -= a
//		}
//	}
//	acc := mylevel.AccumulativePrice
//	//记录差多少越级
//	for i,v := range setlevel{
//		var u upgrade
//		if mylevel.AccumulativePrice < v.LevelvaluePrice && mylevel.AccumulativePrice+orderamount >= v.LevelvaluePrice {
//			u.need = v.LevelvaluePrice-acc
//			u.LevelvaluePrice = mylevel.LevelvaluePrice
//			u.percent = mylevel.Percent
//			if i>0{
//				u.LevelvaluePrice = setlevel[i-1].LevelvaluePrice
//				u.percent = setlevel[i-1].Percentreality
//			}
//			up = append(up,u)
//			acc = v.LevelvaluePrice
//			continue
//		}
//		//最后一个也要放进去
//		if mylevel.AccumulativePrice+orderamount < v.LevelvaluePrice&&len(up)>0{
//			u.LevelvaluePrice = setlevel[i-1].LevelvaluePrice
//			u.percent = setlevel[i-1].Percentreality
//			if mylevel.AccumulativePrice > v.LevelvaluePrice{
//				u.need = orderamount
//			}else{
//				u.need = mylevel.AccumulativePrice+orderamount - setlevel[i-1].LevelvaluePrice
//			}
//			up = append(up,u)
//			break
//		}
//	}
//	isps = append(isps, ip)
//
//	//判断是否存在升级情况
//	if len(up) > 0 {
//		mp := make(map[string]entity.OrderProfit)
//		for _,u := range up{
//			for i,v := range ulevel{
//				if u.LevelvaluePrice < v.LevelvaluePrice{
//					//var isp InvestmentShareProfit
//					isp := mp[v.CustomerId]
//					isp.UserId = v.CustomerId
//					isp.OrderId = param["order_id"]
//					isp.CustomerId = param["userid"]
//					camount := u.need
//					prelevel := u.percent
//					if i>0{
//						prelevel = ulevel[i-1].Percent
//					}
//					for j:=len(setlevel)-1;j>=0;j--{
//						if v.AccumulativePrice + camount > setlevel[j].LevelvaluePrice{
//							//平级
//							if v.AccumulativePrice > setlevel[j].LevelvaluePrice{
//								isp.Profits += camount * (v.Percent-prelevel)
//								break
//							}
//							//越级
//							a :=  v.AccumulativePrice + camount - setlevel[j].LevelvaluePrice
//							isp.Profits += a * (v.Percent-prelevel)
//							camount -= a
//						}
//					}
//					mp[v.CustomerId] = isp
//					//isps = append(isps, isp)
//				}
//			}
//		}
//		for k,_ := range mp{
//			isps = append(isps, mp[k])
//		}
//	}else{
//		for i,v := range ulevel{
//			//if mylevel.Levelvalue < v.Levelvalue{
//			var isp entity.OrderProfit
//			isp.UserId = v.CustomerId
//			isp.CustomerId = param["userid"]
//			isp.OrderId = param["order_id"]
//			camount := orderamount
//			prelevel := mylevel.Percent
//			if i>0{
//				prelevel = ulevel[i-1].Percent
//			}
//			for j:=len(setlevel)-1;j>=0;j--{
//				if v.AccumulativePrice + camount > setlevel[j].LevelvaluePrice{
//					//平级
//					if v.AccumulativePrice > setlevel[j].LevelvaluePrice{
//						isp.Profits += camount * (v.Percent-prelevel)
//						break
//					}
//					//越级
//					a :=  v.AccumulativePrice + camount - setlevel[j].LevelvaluePrice
//					isp.Profits += a * (v.Percent-prelevel)
//					camount -= a
//				}
//			}
//			isps = append(isps, isp)
//			//}
//		}
//	}
//
//	customer := 0.0
//	sql := `insert into ordersprofit(order_id,user_id,profits,customer_id)values`
//	for _,v := range isps{
//		if v.Profits ==0{
//			continue
//		}
//		sql += `(`+v.OrderId+`,`+v.UserId+`,`+utils.Float64ToString(v.Profits)+`,`+v.CustomerId+`),`
//		sql1 := `update customer set balance = balance+ `+utils.Float64ToString(v.Profits) +` where customer_id = `+v.UserId
//		if _,err = session.Exec(sql1);err!=nil{
//			return 0,err
//		}
//		customer += v.Profits
//	}
//	company := orderamount * in.Project.Salesmanratio-customer
//	if company >0{
//		sql += `(`+param["order_id"]+`,0,`+utils.Float64ToString(company)+","+param["userid"]+`)`
//	}else{
//		company = 0
//	}
//
//	sql = strings.TrimRight(sql,",")
//	_,err = session.Exec(sql)
//	return company,err
//}
func (in *OrderProfitModels) AddOrderProfit(session *xorm.Session, param map[string]string, price float64) (float64, error) {
	//var le entity.UserLevel
	//var re Referrer
	//获取订单信息
	//ins ,err := in.GetOrderProfitById(param["investmentid"])
	//if err!=nil{
	//	return err
	//}
	Referrers, err := in.GetReferrers(param["userid"])
	if err != nil {
		return 0, err
	}
	//获取直系上级等级
	mylevel, err := in.GetSetUserByUserid(param["referrer_id"], price)
	if err != nil {
		return 0, err
	}
	//获取上级等级
	//排除掉直系上级 排除0
	//Referrers = strings.Replace(Referrers,param["referrer_id"],"",-1)
	//Referrers = strings.Replace(Referrers,"0","",1)
	refs := strings.Split(Referrers, ",")
	refstr := ""
	for _, v := range refs {
		if v == "0" || v == param["referrer_id"] {
			continue
		}
		refstr += v + ","
	}
	refstr = strings.Trim(refstr, ",")
	var ulevel []entity.UserLevel

	if refstr != "" {
		ulevel, err = in.GetReferrerLevel(refstr, price)
		if err != nil {
			return 0, err
		}
		//ulevel = append(ulevel, le...)
	}

	//获取设置等级
	setlevel, err := in.GetSetUserLevel(mylevel.Levelvalue, price)
	//升级所需
	type upgrade struct {
		need            float64
		LevelvaluePrice float64
		percent         float64
	}
	var up []upgrade
	//计算自己的利润
	var isps []entity.OrderProfit
	var ip entity.OrderProfit
	orderamount, _ := strconv.ParseFloat(param["amount"], 64)
	camount := orderamount
	ip.UserId = param["referrer_id"]
	ip.OrderId = param["order_id"]
	ip.CustomerId = param["userid"]

	for i := len(setlevel) - 1; i >= 0; i-- {
		if mylevel.AccumulativePrice+orderamount > setlevel[i].LevelvaluePrice {
			//平级
			if mylevel.AccumulativePrice > setlevel[i].LevelvaluePrice {
				ip.Profits += camount * setlevel[i].Percentreality
				break
			}
			//越级
			a := mylevel.AccumulativePrice + camount - setlevel[i].LevelvaluePrice
			ip.Profits += a * setlevel[i].Percentreality
			camount -= a
		}
	}
	acc := mylevel.AccumulativePrice
	//记录差多少越级
	for i, v := range setlevel {
		var u upgrade
		if mylevel.AccumulativePrice < v.LevelvaluePrice && mylevel.AccumulativePrice+orderamount >= v.LevelvaluePrice {
			u.need = v.LevelvaluePrice - acc
			u.LevelvaluePrice = mylevel.LevelvaluePrice
			u.percent = mylevel.Percent
			if i > 0 {
				u.LevelvaluePrice = setlevel[i-1].LevelvaluePrice
				u.percent = setlevel[i-1].Percentreality
			}
			up = append(up, u)
			acc = v.LevelvaluePrice
			continue
		}
		//最后一个也要放进去
		if mylevel.AccumulativePrice+orderamount < v.LevelvaluePrice && len(up) > 0 {
			u.LevelvaluePrice = setlevel[i-1].LevelvaluePrice
			u.percent = setlevel[i-1].Percentreality
			if mylevel.AccumulativePrice > v.LevelvaluePrice {
				u.need = orderamount
			} else {
				u.need = mylevel.AccumulativePrice + orderamount - setlevel[i-1].LevelvaluePrice
			}
			up = append(up, u)
			break
		}
	}
	isps = append(isps, ip)

	//判断是否存在升级情况
	if len(up) > 0 {
		mp := make(map[string]entity.OrderProfit)
		for _, u := range up {
			for i, v := range ulevel {
				if u.LevelvaluePrice < v.LevelvaluePrice {
					//var isp InvestmentShareProfit
					isp := mp[v.CustomerId]
					isp.UserId = v.CustomerId
					isp.OrderId = param["order_id"]
					isp.CustomerId = param["userid"]
					camount := u.need
					prelevel := u.percent
					if i > 0 {
						prelevel = ulevel[i-1].Percent
					}
					for j := len(setlevel) - 1; j >= 0; j-- {
						if v.AccumulativePrice+camount > setlevel[j].LevelvaluePrice {
							//平级
							if v.AccumulativePrice > setlevel[j].LevelvaluePrice {
								isp.Profits += camount * (v.Percent - prelevel)
								break
							}
							//越级
							a := v.AccumulativePrice + camount - setlevel[j].LevelvaluePrice
							isp.Profits += a * (v.Percent - prelevel)
							camount -= a
						}
					}
					mp[v.CustomerId] = isp
					//isps = append(isps, isp)
				}
			}
		}
		for k, _ := range mp {
			isps = append(isps, mp[k])
		}
	} else {
		for i, v := range ulevel {
			//if mylevel.Levelvalue < v.Levelvalue{
			var isp entity.OrderProfit
			isp.UserId = v.CustomerId
			isp.CustomerId = param["userid"]
			isp.OrderId = param["order_id"]
			camount := orderamount
			prelevel := mylevel.Percent
			if i > 0 {
				prelevel = ulevel[i-1].Percent
			}
			for j := len(setlevel) - 1; j >= 0; j-- {
				if v.AccumulativePrice+camount > setlevel[j].LevelvaluePrice {
					//平级
					if v.AccumulativePrice > setlevel[j].LevelvaluePrice {
						isp.Profits += camount * (v.Percent - prelevel)
						break
					}
					//越级
					a := v.AccumulativePrice + camount - setlevel[j].LevelvaluePrice
					isp.Profits += a * (v.Percent - prelevel)
					camount -= a
				}
			}
			isps = append(isps, isp)
			//}
		}
	}

	customer := 0.0
	sql := `insert into ordersprofit(order_id,user_id,profits,customer_id)values`
	for _, v := range isps {
		if v.Profits == 0 {
			continue
		}
		//业务部门
		if v.UserId == "0" {
			continue
		}
		sql += `(` + v.OrderId + `,` + v.UserId + `,` + utils.Float64ToString(v.Profits) + `,` + v.CustomerId + `),`
		sql1 := `update customer set balance = balance+ ` + utils.Float64ToString(v.Profits) + ` where customer_id = ` + v.UserId
		if _, err = session.Exec(sql1); err != nil {
			return 0, err
		}
		customer += v.Profits
	}
	company := orderamount*in.Project.Salesmanratio - customer
	if company > 0 {
		sql += `(` + param["order_id"] + `,0,` + utils.Float64ToString(company) + "," + param["userid"] + `)`
	} else {
		company = 0
	}

	sql = strings.TrimRight(sql, ",")
	_, err = session.Exec(sql)
	return company, err
}
func (r *OrderProfitModels) GetReferrers(userid string) (string, error) {
	sql := `select referrers from referrer where userid = ?`
	res, err := r.QueryString(sql, userid)
	if err != nil || len(res) == 0 {
		return "", err
	}
	return res[0]["referrers"], err
}
func (y *OrderProfitModels) GetOrderProfitList(param map[string]string) (interface{}, error) {

	sql := `select profits customer_income,create_time time,id,1 stype 
		FROM ordersprofit where user_id =:user_id `
	sql = utils.SqlReplaceParames(sql, param)
	var err error
	param["total"], err = utils.GetTotalCount(y.Engine, sql)
	if err != nil {
		return nil, err
	}
	param["sort"] = "time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	//var res []entity.SettleLog
	res, err := y.Engine.QueryString(sql)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *OrderProfitModels) GetOrderProfitById(param map[string]string) (map[string]string, error) {

	sql := `select o.order_id,id,profits,o.create_time,1 stype,c.phone FROM ordersprofit o
		left join customer c on c.customer_id = o.customer_id
		where o.user_id =:user_id  and id = :id  `
	sql = utils.SqlReplaceParames(sql, param)

	res, err := y.Engine.QueryString(sql)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return res[0], err
}

func (y *OrderProfitModels) CollectInfo2(param string) (map[string]string, error) {

	sql := `SELECT sum(total_income)total_income,sum(today_income)today_income,	round(sum(accumulative),0)accumulative FROM(
	select sum(profits)total_income,0 today_income,0 accumulative from ordersprofit where user_id = ?
		UNION
	select 0 total_income,sum(profits) today_income,0 accumulative from ordersprofit where date(now()) = date(create_time) and user_id = ?
		UNION
	select 0 total_income,0 total_income,accumulative from customer where customer_id = ?
	)a`
	res, err := y.Engine.QueryString(sql, param, param, param)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return res[0], nil
}

func (this *OrderProfitModels) LowerPowerPerformance(customerID string, param map[string]string) (retList []map[string]string, err error) {
	// 先取出自己的直接下级
	sql := `select customer_id from customer where referrer_id=%s`
	sql = fmt.Sprintf(sql, customerID)
	param["total"], err = utils.GetTotalCount(this.Engine, sql)
	if err != nil {
		return
	}

	pageSize, err := strconv.ParseInt(param["pageSize"], 10, 64)
	if err != nil {
		return
	}
	pageIndex, err := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err != nil {
		return
	}
	start := (pageIndex - 1) * pageSize
	sql += " limit " + strconv.FormatInt(start, 10) + "," + param["pageSize"]

	retMapList, err := this.Engine.QueryString(sql)
	if err != nil {
		return
	}

	retList = make([]map[string]string, 0)
	for _, mapData := range retMapList {
		data, err1 := this.getPerformance(mapData["customer_id"], customerID)
		if err1 != nil {
			return
		}
		retList = append(retList, data)
	}

	return
}

func (this *OrderProfitModels) getPerformance(customerID, profitCustomerID string) (mapData map[string]string, err error) {
	mapData = make(map[string]string)
	sql := `select c.phone, ifnull(a.total, 0) performance from customer c 
	left join (select sum(p.hashrate) total, p.customer_id from profit_record p where p.customer_id=%s and p.profit_customer_id=%s) a on a.customer_id = c.customer_id
  		where c.customer_id = %s`
	sql = fmt.Sprintf(sql, customerID, profitCustomerID, customerID)
	mapList, err := this.Engine.QueryString(sql)
	if err != nil {
		return
	}
	if len(mapList) <= 0 {
		log.Error("没找到客户数据，客户ID:", profitCustomerID)
		return
	}
	mapData = mapList[0]
	return
}

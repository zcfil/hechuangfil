package models

import (
	"fmt"
	"hechuangfil/utils"
	"strings"
)

//下单
func (y *UserModels) PlaceTheOrder(param map[string]string)error{
	//获取商品库存
	sess := y.Engine.NewSession()
	var err error
	defer func() {
		if err!=nil{
			sess.Rollback()
			return
		}
		sess.Commit()
	}()
	sess.Begin()
	//扣掉余额
	sql1 := `update customer set balance = balance - :amount where customer_id=:userid `

	sql1 = utils.SqlReplaceParames(sql1,param)
	if _,err = sess.Exec(sql1);err!=nil{
		return err
	}
	price,_ := y.GetHashratePrice()
	var company float64
	if company,err = y.AddOrderProfit(sess,param,price);err!=nil{
		return err
	}
	param["company"] = utils.Float64ToString(company)
	sql := "insert into orders(order_id,amount,customer_id,remark,hashrate,expiration_date,salesdep_profit)\n  " +
		"value (:order_id,:amount,:userid,:remark,:hashrate,ADDDATE(now(),INTERVAL 540 day),:company)"
	sql = utils.SqlReplaceParames(sql,param)

	_, err = sess.Exec(sql)
	if err!=nil{
		return err
	}

	sql2 := ` update customer set accumulative = accumulative+`+param["hashrate"]+` where customer_id =`+param["userid"]

	if _,err = sess.Exec(sql2);err!=nil{
		return err
	}

	setlevel,_ := y.GetSetUserLevel(0,price)
	if len(setlevel)>0{
		//acc := y.GetAccumulative(param["referrer_id"])
		//if mylevel[0].Accumulative<=acc{
		//	sql3 := ` update customer set accumulative = accumulative+`+param["hashrate"]+` where customer_id =`+param["referrer_id"]
		//	if _,err = sess.Exec(sql3);err!=nil{
		//		return err
		//	}
		//}
		ref,err := y.GetReferrers(param["userid"])
		if err!=nil{
			return err
		}
		refs := strings.Split(ref,",")
		for _,v := range refs{
			acc := y.GetAccumulative(v)
			if setlevel[0].Levelvalue<=acc{
				sql3 := ` update customer set accumulative = accumulative+`+param["hashrate"]+` where customer_id =`+v
				if _,err = sess.Exec(sql3);err!=nil{
					return err
				}

				// 记录给自己上级添加业绩
				sql4 := `insert into profit_record (customer_id, hashrate, amount, profit_customer_id, order_id, date) values (%s, %s, %s, %s, %s, now())`
				sql4 = fmt.Sprintf(sql4, param["userid"],param["hashrate"],param["amount"], v, param["order_id"])
				if _, err = sess.Exec(sql4); err != nil {
					return err
				}
			}
		}
	}

	return err
}

//订单列表
func (y *UserModels) GetOrderList (param map[string]string)([]map[string]string,map[string]string,error){
	sql := ` select * from orders where customer_id = ? `
	param["sort"] = "create_time"
	param["order"] = "desc"
	//param["total"],_ = utils.GetTotalCount(y.Engine,sql,param["user_id"])
	sqltotal := `select sum(hashrate)hashrate,count(1) total from(`+sql+`)a`
	restotal,err := y.Engine.QueryString(sqltotal,param["user_id"])
	if err!=nil{
		return restotal,nil,err
	}
	param["total"] = restotal[0]["total"]
	total := make(map[string]string)
	total["hashrate"] = restotal[0]["hashrate"]
	sql += utils.LimitAndOrderBy(param)

	res,err := y.Engine.QueryString(sql,param["user_id"])
	if err!=nil{
		return res,nil,err
	}

	return res,total,err
}

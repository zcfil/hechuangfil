package models

import (
	"errors"
	"github.com/go-xorm/xorm"
	"hechuangfil/utils"
)

func (y *UserModels) RechargeCoinFil(param map[string]string)error{

	sql := "insert into recharge(amount,customer_id,salesman_id)\n  " +
		"value (:amount,:userid,:salesman_id)"

	sql = utils.SqlReplaceParames(sql,param)

	if _,err := y.Engine.Exec(sql);err!=nil{
		return err
	}
	return nil
}

func (y *UserModels) WithdrawCoinFil(param map[string]string)error{
	sess := y.NewSession()
	sess.Begin()
	var err error
	defer func() {
		if err!=nil{
			sess.Rollback()
			return
		}
		sess.Commit()
	}()
	sql1 := "update customer set frozen_capital = frozen_capital+:amount,balance =balance-:amount where customer_id=:userid "
	sql1 = utils.SqlReplaceParames(sql1,param)
	if _,err = sess.Exec(sql1);err!=nil{
		return err
	}

	sql := "insert into withdraw(amount,to_addres,customer_id)\n  " +
		"value (:amount,:to_address,:userid)"

	sql = utils.SqlReplaceParames(sql,param)

	if _,err = sess.Exec(sql);err!=nil{
		return err
	}
	return nil
}

func (y *UserModels) WalletBalance(userid string)(map[string]string,error){

	sql := "select wallet,balance,frozen_capital from customer where customer_id = ? "

	res,err := y.Engine.QueryString(sql,userid)
	if err!=nil||len(res)==0{
		return nil, err
	}
	return res[0],nil
}

func (y *UserModels) RechargeList(param map[string]string) ( res []map[string]string, err error) {

	sql := `select * from recharge where customer_id =?  `

	param["total"], err = utils.GetTotalCount(y.Engine, sql,param["user_id"])
	if err != nil {
		return
	}
	param["sort"] = "create_time"
	param["order"] = "desc"
 	sql += utils.LimitAndOrderBy(param)
	res, err = y.Engine.QueryString(sql,param["user_id"])
	if err != nil {
		return
	}
	return
}
func (y *UserModels) RechargeById(param map[string]string) ( map[string]string, error) {

	sql := `select * from recharge where customer_id =:user_id and recharge_id = :recharge_id `
	sql = utils.SqlReplaceParames(sql,param)

	res, err := y.Engine.QueryString(sql)
	if err != nil|| len(res)==0 {
		return nil,err
	}
	return res[0],err
}

func (y *UserModels) WithdrawList(param map[string]string) ( res []map[string]string, err error) {

	sql := `select * from withdraw where customer_id =? `

	param["total"], err = utils.GetTotalCount(y.Engine, sql,param["user_id"])
	if err != nil {
		return
	}
	param["sort"] = "create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	res, err = y.Engine.QueryString(sql,param["user_id"])
	if err != nil {
		return
	}
	return
}
func (y *UserModels) WithdrawById(param map[string]string) ( map[string]string, error) {

	sql := `select * from withdraw where customer_id =:user_id and withdraw_id = :withdraw_id `
	sql = utils.SqlReplaceParames(sql,param)
	res, err := y.Engine.QueryString(sql)
	if err != nil|| len(res)==0 {
		return nil,err
	}
	return res[0],err
}

func (y *UserModels) CollectInfo(param string) ( map[string]string, error) {

	sql := `select sum(total_income)total_income,sum(to_customer_balance)to_customer_balance,sum(to_customer_lock)to_customer_lock,sum(today_income)today_income from(
		select sum(to_customer_balance+to_customer_lock)total_income,sum(to_customer_balance+customer_lock_release)to_customer_balance
					,sum(to_customer_lock) to_customer_lock,0 today_income from settle_log where customer_id = ?
		union
		select 0 total_income,0 to_customer_balance,0 to_customer_lock,sum(to_customer_balance+to_customer_lock)today_income 
					from settle_log 
					where customer_id = ? and date(time) = date(now())
		)a `
	res, err := y.Engine.QueryString(sql,param,param)
	if err != nil|| len(res)==0 {
		return nil,err
	}
	return res[0],nil
}

// 获取下线列表
func (y *UserModels) GetLowerLevelPerformance() {

}

//修改提币钱包
func (y *UserModels) SetWithdrawWallet(sess *xorm.Session,param map[string]string) ( error) {
	if utils.CheckAddress(param["wallet"]){
		return errors.New("提币钱包格式不正确！")
	}
	sql := `update customer set withdraw_wallet = :wallet where customer_id = :user_id `
	sql = utils.SqlReplaceParames(sql,param)
	_, err := sess.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (y *UserModels) GetSettleLogList(param map[string]string) ( interface{}, error) {

	sql := `select customer_income,time,id,0 stype from settle_log 
		where customer_id =:user_id `
	sql = utils.SqlReplaceParames(sql,param)
	var err error
	param["total"], err = utils.GetTotalCount(y.Engine,sql)
	if err!=nil{
		return nil, err
	}
	param["sort"] = "time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	//var res []entity.SettleLog
	res,err := y.Engine.QueryString(sql)
	if err != nil {
		return nil,err
	}
	return res,nil
}

func (y *UserModels) GetSettleLogById(param map[string]string) ( map[string]string, error) {

	sql := `select customer_income,to_customer_balance,to_customer_lock,customer_lock_release,id,time from settle_log 
				where id = :id and customer_id = :user_id `
	sql = utils.SqlReplaceParames(sql,param)

	res, err := y.Engine.QueryString(sql)
	if err != nil|| len(res)==0 {
		return nil,err
	}
	return res[0],err
}



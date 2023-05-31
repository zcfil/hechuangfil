package models

import (
	"github.com/go-xorm/xorm"
	"hechuangfil/entity"
	"strconv"
	"strings"
)

type UserLevelModels struct {
	*xorm.Engine
}
func NewUserLevelModels(x *xorm.Engine)*UserLevelModels{
	return &UserLevelModels{
		x,
	}
}
func (u *UserLevelModels)mapToUserLevel(mp []map[string]string,price float64)[]entity.UserLevel{
	var ul []entity.UserLevel
	for _,v := range mp{
		var u entity.UserLevel
		u.Levelvalue,_ = strconv.ParseFloat(v["levelvalue"],64)
		u.LevelvaluePrice = u.Levelvalue * price			//算力转实际设置的FIL

		u.Accumulative,_ = strconv.ParseFloat(v["accumulative"],64)
		u.AccumulativePrice = u.Accumulative * price		//算力转实际设置的FIL

		u.Percent,_ = strconv.ParseFloat(v["percent"],64)
		u.Percentreality,_ = strconv.ParseFloat(v["percentreality"],64)
		u.CustomerId,_ = v["customer_id"]
		ul = append(ul, u)
	}
	return ul
}
func (u *UserLevelModels)GetUserLevelList(ids string,price float64)([]entity.UserLevel,error){
	con := `and customer_id in (`+ids+`)`
	sql := `select u.customer_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from customer u
					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
					where 1=1 `+con+`
					GROUP BY customer_id order by levelvalue`

	res,err := u.Engine.QueryString(sql)
	if  err!=nil{
		return nil,err
	}

	return u.mapToUserLevel(res,price),nil
}
//排除一样等级的
func (u *UserLevelModels)GetUserLevel(ids string,levelvalue float64,price float64)([]entity.UserLevel,error){
	con := `and customer_id in (`+ids+`)`
	sql := `select u.customer_id,u.accumulative, b.levelvalue, b.percent from customer u
					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
					left join (
							select count(1) count,levelvalue,sum(percent) percent from (
							select u.customer_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
							left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
							where 1=1 `+con+` and l.levelvalue > ?
							GROUP BY customer_id order by levelvalue
							)a GROUP BY levelvalue 
					)b on b.levelvalue = l.levelvalue
					where b.count = 1 `+con+` and l.levelvalue > ?
					order by levelvalue`
	res,err := u.QueryString(sql,levelvalue,levelvalue)
	if err!=nil{
		return nil,err
	}

	return u.mapToUserLevel(res,price),err
}
func (u *UserLevelModels)GetReferrerLevel(ids string,price float64)([]entity.UserLevel,error){
	sql := `select u.customer_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from customer u
					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
					where customer_id in (`+ids+`) and is_del = 0 and status = 0
					GROUP BY customer_id order by levelvalue`
	re,err := u.QueryString(sql)
	//mp := make(map[string]UserLevel)
	//for i:=0;i<len(ul);i++ {
	//	mp[ul[i].UserId] = ul[i]
	//}
	//同级排序，直接上级在前面
	refs := strings.Split(ids,",")
	mp := make(map[string]int)
	for i := len(refs)-1;i>=0;i--{
		mp[refs[i]] = i
	}
	ul := u.mapToUserLevel(re,price)
	var res []entity.UserLevel
	//一级只取一个
	for i:=0;i<len(ul);i++{
		pre := ul[i]
		for j:=i+1;j<len(ul);j++{
			if pre.Levelvalue < ul[j].Levelvalue{
				//res = append(res, ul[i])
				break
			}else{
				i++
			}
			if mp[pre.CustomerId]<mp[ul[j].CustomerId]{
				pre = ul[j]
			}
		}
		res = append(res, pre)
	}
	return res,err
}

//获取设置等级
func (u *UserLevelModels)GetSetUserLevel(levelvalue float64,price float64)([]entity.UserLevel,error){
	sql := `select * from user_level where levelvalue>=? order by levelvalue `

	res,err := u.QueryString(sql,levelvalue)

	return u.mapToUserLevel(res,price),err
}

//获取自己等级
func (u *UserLevelModels)GetSetUserByUserid(userid string,price float64)(entity.UserLevel,error){
	sql := `select u.customer_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent, max(l.percentreality) percentreality
			from customer u
			left join (select * from user_level ) l on  u.accumulative >= l.levelvalue
			where u.customer_id = `+userid+`
			GROUP BY customer_id 
			order by levelvalue `
	re,err := u.QueryString(sql)
	if err!=nil || len(re)==0{
		return entity.UserLevel{},err
	}
	return u.mapToUserLevel(re,price)[0],nil
}


package models

import (
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"go.uber.org/dig"
	"hechuangfil/conf"
	"hechuangfil/service"
	"strconv"
)

type LoginModels struct {
	dig.In
	*xorm.Engine
	*redis.Client
	*service.Response
	*conf.AuthCode
}

func (this *LoginModels) CanRegisterByPhone(phone string) bool {
	sql := `select customer_id from customer where phone = "` + phone + `" and is_del = 0 limit 1`
	res, err := this.Engine.QueryString(sql)
	if err != nil {
		return false
	}
	if len(res) > 0 {
		return false
	}
	return true
}

func (l *LoginModels)GetUserPwd(phone string)string{
	sql := `select password from user where phone='`+phone+ `'`
	res,_ := l.Engine.QueryString(sql)
	return res[0]["password"]
}

func (l *LoginModels)SetReferrer(sess *xorm.Session, userid int64,referrerid string)error{
	sql := `select * from referrer where userid=?`
	ref,err := l.Engine.QueryString(sql,referrerid)
	if err!=nil{
		return err
	}
	referrers := ""
	if len(ref)>0{
		referrers = ref[0]["referrers"]
	}
	if(len(referrers)>0){
		referrers += ","
	}
	referrers += referrerid
	if referrerid!="0"{
		sql2 := `update referrer set referrals = CONCAT(ifnull(referrals,''),',`+strconv.FormatInt(userid,10)+`') where userid in (`+referrers+`) `
		if _,err = sess.Exec(sql2);err!=nil{
			return err
		}
	}

	sql3 := `insert into referrer (userid,referrers)value(`+strconv.FormatInt(userid,10)+`,'`+referrers+`') `
	//e.Referrers = ref.Referrers
	if _,err = sess.Exec(sql3);err!=nil{
		return err
	}

	return err
}

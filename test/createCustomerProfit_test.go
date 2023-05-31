package test

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-xorm/xorm"
	"hechuangfil/conf"
	"hechuangfil/db"
	log "hechuangfil/logrus"
	"strconv"
	"strings"
	"testing"
	"time"
)

// 慎用！慎用！慎用！
// 本测试模块是用来创建历史记录的 慎用！慎用！慎用！


func Test_createCustomerProfitData(t *testing.T) {
	var config conf.Config
	if _, err := toml.DecodeFile("./conf/config.toml", &config); err != nil {
		fmt.Println("Decode config file error:", err.Error())
	}
	fmt.Println("config:", config)

	log.NewLogger(config.LogFile.FileName)

	engine, err := db.NewSqlEngine(config.Mysql)()
	if err != nil {
		return
	}

	sql := `select * from orders`
	mapList, err := engine.QueryString(sql)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for _, m := range mapList {
		sql1 := `select referrers from referrer where userid = %s`
		sql1 = fmt.Sprintf(sql1, m["customer_id"])
		refMapList, err := engine.QueryString(sql1)
		if err == xorm.ErrTableNotFound {
			continue
		}
		if err != nil {
			log.Error("err:", err.Error())
			return
		}

		if len(refMapList) <= 0 {
			continue
		}
		refCustomers := strings.Split(refMapList[0]["referrers"], ",")
		for _, refCustomerID := range refCustomers {
			if refCustomerID == "0" {
				continue
			}

			// 需要判断获益者的时间
			if !canInsert(refCustomerID, engine, m["create_time"]) {
				continue
			}
			sql2 := `insert into profit_record (customer_id, hashrate, amount, profit_customer_id, order_id, date) values (%s, %s, %s, %s, %s, '%s')`
			sql2 = fmt.Sprintf(sql2, m["customer_id"], m["hashrate"], m["amount"], refCustomerID, m["order_id"], m["create_time"])
			if _, err = engine.Exec(sql2); err != nil {
				log.Error("err:", err.Error())
			}


		}
	}
}

func canInsert(refCustomerID string, engine *xorm.Engine, orderTime string) bool {
	sql := `select * from orders where customer_id = %s  order by create_time asc`
	sql = fmt.Sprintf(sql, refCustomerID)
	mapList, err := engine.QueryString(sql)
	if err != nil {
		log.Error("err:", err.Error())
		return false
	}

	hashRate := float64(0)
	lastTime := ""
	for _, m := range mapList {
		hr, err := strconv.ParseFloat(m["hashrate"], 64)
		if err != nil {
			log.Error("err:", err.Error())
			continue
		}
		hashRate += hr
		if hashRate >= 3 {
			lastTime = m["create_time"]
			break
		}
	}

	if lastTime == "" {
		return false
	}

	if dateStr2Time(orderTime).Unix() > dateStr2Time(lastTime).Unix() {
		return true
	}
	return false
}

func dateStr2Time(datetime string) time.Time {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	return tmp
}

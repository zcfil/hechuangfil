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
)

func Test_selectData(t *testing.T) {
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

	list1 := make([]string, 0)
	list2 := make([]string, 0)
	mapReachTime := make(map[string]string)
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

		//fmt.Println(refMapList)

		if len(refMapList) <= 0 {
			continue
		}
		refCustomers := strings.Split(refMapList[0]["referrers"], ",")
		for _, refCustomerID := range refCustomers {
			if refCustomerID == "0" {
				continue
			}
			dateStr, err := reachDate(refCustomerID, engine)
			if err != nil {
				continue
			}

			if dateStr == "" {
				continue
			}

			if _, ok := mapReachTime[refCustomerID]; ok {
				continue
			}

			mapReachTime[refCustomerID] = dateStr


			//list1 = append(list1, fmt.Sprintln("customer:1", m["customer_id"], " refCustomer:", refCustomerID))
			//list2 = append(list2, fmt.Sprintln("customer:1", m["customer_id"], " refCustomer:", refCustomerID))
			//	sql2 := `insert into profit_record (customer_id, hashrate, amount, profit_customer_id, order_id, date) values (%s, %s, %s, %s, %s, '%s')`
			//	sql2 = fmt.Sprintf(sql2, m["customer_id"], m["hashrate"], m["amount"], refCustomerID, m["order_id"], m["create_time"])
			//	if _, err = engine.Exec(sql2); err != nil {
			//		log.Error("err:", err.Error())
			//	}
			//
			//
		}
	}

	fmt.Println("list1:", list1)
	fmt.Println("list2:", list2)
	fmt.Println("mapDate:", mapReachTime)
}

func reachDate(customerID string, engine *xorm.Engine) (timeStr string, err error) {
	sql := `select * from orders where customer_id = %s  order by create_time asc`
	sql = fmt.Sprintf(sql, customerID)
	mapList, err := engine.QueryString(sql)
	if err != nil {
		log.Error("err:", err.Error())
		return
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

	timeStr = lastTime
	return
}
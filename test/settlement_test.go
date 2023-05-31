package test

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"hechuangfil/conf"
	"hechuangfil/db"
	log "hechuangfil/logrus"
	"math/rand"
	"strconv"
	"testing"
)


func Test_insertOrder(t *testing.T) {
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

	//userCount := 10

	//idList := make([]int64, 0)
	//for i := 0; i < userCount; i++ {
	//	var c entity.Customer
	//	c.Phone = "1377777777" + strconv.FormatInt(int64(i), 10)
	//	engine.Insert(c)
	//
	//}

	idList := []int64{1432250014607872001,1432250014607872002,1432250014607872003,1432250014607872004,1432250014607872005}
	for i :=0; i< 100; i++ {
		customerID := idList[rand.Intn(5)]
		amount := rand.Intn(10000)
		orderID := int64(uuid.New().ID())
		sql := `insert into orders(order_id,amount,customer_id,salesman_id,hashrate) value (` +
				strconv.FormatInt(orderID, 10) + `,` + strconv.FormatInt(int64(amount), 10) + `,` + strconv.FormatInt(customerID, 10) + `, 0, 7)`
		engine.Exec(sql)
	}


}

func Test_settlement(t *testing.T) {
	var config conf.Config
	if _, err := toml.DecodeFile("./conf/config.toml", &config); err != nil {
		fmt.Println("Decode config file error:", err.Error())
	}
	fmt.Println("config:", config)

	log.NewLogger(config.LogFile.FileName)

	//var settle models.Settlement
	//engine, err := db.NewSqlEngine(config.Mysql)()
	//if err != nil {
	//	return
	//}
	//
	//settle.Engine = engine

}

func Test_checkTime(t *testing.T) {
	var config conf.Config
	if _, err := toml.DecodeFile("./conf/config.toml", &config); err != nil {
		fmt.Println("Decode config file error:", err.Error())
	}
	fmt.Println("config:", config)

	log.NewLogger(config.LogFile.FileName)

	//var settle models.Settlement
	//engine, err := db.NewSqlEngine(config.Mysql)()
	//if err != nil {
	//	return
	//}
	//
	//settle.Engine = engine
	//settle.SettleFunc()
}
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"hechuangfil/conf"
)

//NewSqlEngine xorm引擎配置
func NewSqlEngine(mysql *conf.MySql) func() (*xorm.Engine, error) {

	return func() (*xorm.Engine, error){
		engine, e := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8",mysql.Username,mysql.Password,mysql.Host,mysql.Name))
		engine.ShowSQL(true)
		return engine, e
	}

}

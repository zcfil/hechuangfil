package conf

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter/v2"
	"log"
)

//加载 yungo权限
func LoadCasbin(mysql *MySql) func() (*casbin.Enforcer, error) {

	return func() (*casbin.Enforcer, error) {

		//从db加载policy
		a, err := xormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",mysql.Username,mysql.Password,mysql.Host,mysql.Name), true)
		if err != nil {
			log.Printf("连接数据库错误: %v", err)
			return nil, err
		}
		en, err := casbin.NewEnforcer("./conf/model.conf", a)
		// Load the policy from DB.
		en.LoadPolicy()
		return en, err
	}

}

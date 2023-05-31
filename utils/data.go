package utils

import (
	"strconv"
	"strings"
	"github.com/go-xorm/xorm"
	"hechuangfil/entity"
)

func DataToFloat64(mp map[string]string,key string)float64{
	re,_ := strconv.ParseFloat(mp[key],64)
	return re
}

func DataToInt64(mp map[string]string,key string)int64{
	re,_ := strconv.ParseInt(mp[key],10,64)
	return re
}

func DataToString(mp []map[string][]byte)*[]map[string]string{
	var re = make([]map[string]string,len(mp))
	for index,v := range mp{
		m := make(map[string]string)
		for k1,v1 := range v{
			m[k1] = string(v1)
		}
		re[index] = m
	}
	return &re
}

func GetTotalCount(dbtype *xorm.Engine, sql string, value ...string) (string,error) {
	sql = ` select count(1) total from ( ` + sql + `)a`
	sqltotal := ""
	if strings.Contains(sql, "?") {
		str := strings.Split(sql, "?")
		for i := 0; i < len(str); i++ {
			if i < len(str)-1 {
				sqltotal += str[i] + `'` + value[i] + `'`
			} else {
				sqltotal += str[i]
			}
		}
	} else {
		sqltotal = sql
	}
	res,err := dbtype.QueryString(sqltotal)

	return res[0]["total"],err
}
func GetConfig(dbtype *xorm.Engine,c_name string,user *entity.SysUser) (string) {
	sql := ` select c_name,c_value from sys_config where c_name = :c_name `
	param := make(map[string]string)
	param["c_name"] = c_name
	sql = SqlReplaceParames(sql,param)
	res,_ := dbtype.QueryString(sql)
	if len(res)==0{
		return ""
	}
	return res[0]["c_value"]
}
func DataToKeyValue(data []map[string]string,key,value string)map[string]string{
	result := make(map[string]string)
	for _,val := range data{
		keystr := ""
		valuestr := ""
		for k,v := range val{
			if k == key{
				keystr = v
			}
			if k == value{
				valuestr = v
			}
		}
		result[keystr] = valuestr
	}

	return result
}
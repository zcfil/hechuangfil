package service

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"github.com/tealeg/xlsx"
	"go.uber.org/dig"
	"mime/multipart"
	"strconv"
	"hechuangfil/entity"
	"hechuangfil/utils"
)

type FinanceService struct {
	dig.In
	*redis.Client
	*xorm.Engine
}
func (f *FinanceService)GetDics(str string)map[string]string{
	q1 := ` select dict_name,dict_value from sys_dict where dict_type ='`+str+`' `
	dics,_ := f.Engine.QueryString(q1)
	dickey := make(map[string]string)
	for _,val := range dics{
		key := ""
		value := ""
		for k,v := range val{
			if k == "dict_name"{
				key = v
			}
			if k == "dict_value"{
				value = v
			}
		}
		dickey[key] = value
	}
	return dickey
}
//导入
func (f *FinanceService)UploadApply(file multipart.File,Size int64,user *entity.SysUser)error{
	buf := make([]byte,Size)
	n,_ := file.Read(buf)

	xf ,_ := xlsx.OpenBinary(buf[:n])
	//获取字典
	dickey := f.GetDics("audit")
	statuskey := f.GetDics("status")
	//获取配置
	param := make(map[string]string)
	param["username"] = user.Username
	param["c_name"] = "charge"

	//charge,_ := strconv.ParseFloat(config,64)

	str := ",(:admin_id,(select count(1)+1 a from batchnum a),"
	sql := `insert into withdraw(admin_id,batch_count,`

	for _,sheet := range xf.Sheets{
		exist := make([]string,len(sheet.Rows))
		for j,row := range sheet.Rows{
			value := row.Cells[1].String()
			if value == "" || value == "合计"{
				sql = sql[0:len(sql)-len(str)]
				break
			}
			for i,cell := range row.Cells{
				if j==0{
					if _,ok := dickey[cell.String()];ok{
						exist[i] = cell.String()
						sql += dickey[cell.String()]+","
					}
					if i == len(row.Cells)-1{
						sql = sql[0:len(sql)-1]
						sql += ")values(:admin_id,(select count(1)+1 a from batchnum a where admin_id=:admin_id),"
					}
				}else{
					if len(exist[i])>0{
						value := cell.String()
						if exist[i] == "状态"{
							value = statuskey[value]
						}
						sql += "'"+value+"',"
					}
					if i == len(row.Cells)-1{
						sql = sql[0:len(sql)-1]
						sql += ")"
					}
				}
			}

			if j >= 1 {
				sql += str
			}
		}
	}

	sess:= f.Engine.NewSession()
	sess.Begin()

	sql = utils.SqlReplaceParames(sql,param)
	fmt.Println(sql)
	re,err := f.Engine.Exec(sql)
	if err!=nil{
		fmt.Println("导入错误：",err)
		return err
	}
	count,_ := re.RowsAffected()
	bsql := ` insert into batchnum(operator, num, admin_id,count,create_time)value(:username,(select count(1)+1 a from batchnum a where admin_id=:admin_id) ,:admin_id,:count,now()); `
	if err!=nil{
		sess.Rollback()
		return err
	}
	param["count"] = strconv.FormatInt(count,10)
	bsql = utils.SqlReplaceParames(bsql,param)

	_,err = f.Engine.Exec(bsql)
	if err!=nil{
		sess.Rollback()
		return err
	}

	return sess.Commit()
}
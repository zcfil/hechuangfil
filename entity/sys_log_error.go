package entity

import (
	"time"
)

type SysLogError struct {
	Id            int64     `xorm:"pk comment('id') BIGINT(20)"`
	Module        string    `xorm:"comment('模块名称，如：sys') index VARCHAR(50)"`
	RequestUri    string    `xorm:"comment('请求URI') VARCHAR(200)"`
	RequestMethod string    `xorm:"comment('请求方式') VARCHAR(20)"`
	RequestParams string    `xorm:"comment('请求参数') TEXT"`
	UserAgent     string    `xorm:"comment('用户代理') VARCHAR(500)"`
	Ip            string    `xorm:"comment('操作IP') VARCHAR(32)"`
	ErrorInfo     string    `xorm:"comment('异常信息') TEXT"`
	Creator       int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate    time.Time `xorm:"comment('创建时间') index DATETIME"`
}

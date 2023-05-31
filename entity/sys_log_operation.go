package entity

import (
	"time"
)

type SysLogOperation struct {
	Id            int64     `xorm:"pk comment('id') BIGINT(20)"`
	Module        string    `xorm:"comment('模块名称，如：sys') index VARCHAR(32)"`
	Operation     string    `xorm:"comment('用户操作') VARCHAR(50)"`
	RequestUri    string    `xorm:"comment('请求URI') VARCHAR(200)"`
	RequestMethod string    `xorm:"comment('请求方式') VARCHAR(20)"`
	RequestParams string    `xorm:"comment('请求参数') TEXT"`
	RequestTime   int       `xorm:"not null comment('请求时长(毫秒)') INT(10)"`
	UserAgent     string    `xorm:"comment('用户代理') VARCHAR(500)"`
	Ip            string    `xorm:"comment('操作IP') VARCHAR(32)"`
	Status        int       `xorm:"not null comment('状态  0：失败   1：成功') TINYINT(4)"`
	CreatorName   string    `xorm:"comment('用户名') VARCHAR(50)"`
	Creator       int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate    time.Time `xorm:"comment('创建时间') index DATETIME"`
}

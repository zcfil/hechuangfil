package entity

import (
	"time"
)

type SysLogLogin struct {
	Id          int64     `xorm:"pk comment('id') BIGINT(20)"`
	Operation   int       `xorm:"comment('用户操作   0：用户登录   1：用户退出') TINYINT(3)"`
	Status      int       `xorm:"not null comment('状态  0：失败    1：成功    2：账号已锁定') index TINYINT(3)"`
	UserAgent   string    `xorm:"comment('用户代理') VARCHAR(500)"`
	Ip          string    `xorm:"comment('操作IP') VARCHAR(32)"`
	CreatorName string    `xorm:"comment('用户名') VARCHAR(50)"`
	Creator     int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate  time.Time `xorm:"comment('创建时间') index DATETIME"`
}

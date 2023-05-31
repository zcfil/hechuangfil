package entity

import (
	"time"
)

type YungoUserRole struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	UserId     int64     `xorm:"not null BIGINT(20)"`
	RoleId     int       `xorm:"not null INT(10)"`
	CreateTime time.Time `xorm:"DATETIME"`
	UpdateTime time.Time `xorm:"DATETIME"`
	Creator    string    `xorm:"VARCHAR(20)"`
}

package entity

import (
	"time"
)

type YungoRole struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Role       string    `xorm:"not null unique VARCHAR(20)"`
	Creator    string    `xorm:"VARCHAR(20)"`
	CreateTime time.Time `xorm:"DATETIME"`
	UpdateTime time.Time `xorm:"DATETIME"`
}

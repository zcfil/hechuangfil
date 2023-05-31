package entity

import (
	"time"
)

type YungoMiner struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	MemberId   int64     `xorm:"not null BIGINT(20)"`
	Miner      string    `xorm:"not null VARCHAR(20)"`
	Isdefault  int       `xorm:"not null default 1 comment('1:系统，0:个人') TINYINT(1)"`
	Creator    int64     `xorm:"not null BIGINT(20)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}

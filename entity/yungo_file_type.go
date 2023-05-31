package entity

import (
	"time"
)

type YungoFileType struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	TypeNameZh string    `xorm:"not null VARCHAR(20)"`
	TypeNameEn string    `xorm:"VARCHAR(20)"`
	Creator    string    `xorm:"VARCHAR(20)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Updater    string    `xorm:"VARCHAR(20)"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}

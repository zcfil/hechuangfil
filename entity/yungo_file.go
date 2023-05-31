package entity

import (
	"time"
)

type YungoFile struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	MemberId   int64     `json:"MemberID,string" xorm:"BIGINT(19)"`
	Type       int       `xorm:"not null INT(10)"`
	Filename   string    `xorm:"VARCHAR(200)"`
	FileSize   string    `xorm:"VARCHAR(200)"`
	Price      string    `xorm:"VARCHAR(20)"`
	RootCid    string    `xorm:"VARCHAR(200)"`
	MinerId    string    `xorm:"VARCHAR(20)"`
	DealId     string    `xorm:"VARCHAR(200)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}


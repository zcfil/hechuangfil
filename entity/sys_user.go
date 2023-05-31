package entity

import (
	"time"
)

type SysUser struct {
	Id         int64     `json:"userId,string" xorm:"pk comment('id') BIGINT(20)"`
	Username   string    `json:"username" xorm:"comment('用户名') unique VARCHAR(50)"`
	Password   string    `json:"password" xorm:"comment('密码') VARCHAR(100)"`
	Rank	   string    `json:"rank" xorm:"comment('等级') VARCHAR(10)"`
	SuperAdmin int       `json:"superAdmin" xorm:"comment('超级管理员   0：否   1：是') TINYINT(3)"`
	Delflag    int       `json:"delFlag" xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	CreateDate time.Time `json:"createDate" xorm:"comment('创建时间') index DATETIME"`
}

type UserDept struct {
	SysUser `xorm:"extends"json:"sys_user"`
	SysDept `xorm:"extends" json:"sys_Dept"`
}



package entity

import (
	"time"
)

type SysMenu struct {
	Id          int64     `xorm:"pk comment('id') BIGINT(20)"`
	Pid         int64     `xorm:"comment('上级ID，一级菜单为0') index BIGINT(20)"`
	Url         string    `xorm:"comment('菜单URL') VARCHAR(200)"`
	Type        int       `xorm:"comment('类型   0：菜单   1：按钮') TINYINT(3)"`
	Icon        string    `xorm:"comment('菜单图标') VARCHAR(50)"`
	Permissions string    `xorm:"comment('权限标识，如：sys:menu:save') VARCHAR(32)"`
	Sort        int       `xorm:"comment('排序') INT(11)"`
	DelFlag     int       `xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	Creator     int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate  time.Time `xorm:"comment('创建时间') index DATETIME"`
	Updater     int64     `xorm:"comment('更新者') BIGINT(20)"`
	UpdateDate  time.Time `xorm:"comment('更新时间') DATETIME"`
}

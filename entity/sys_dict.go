package entity

import (
	"time"
)

type SysDict struct {
	Id         int64     `xorm:"pk comment('id') BIGINT(20)"`
	Pid        int64     `xorm:"comment('上级ID，一级为0') index BIGINT(20)"`
	DictType   string    `xorm:"not null comment('字典类型') index VARCHAR(50)"`
	DictName   string    `xorm:"not null comment('字典名称') VARCHAR(255)"`
	DictValue  string    `xorm:"comment('字典值') VARCHAR(255)"`
	Remark     string    `xorm:"comment('备注') VARCHAR(255)"`
	Sort       int       `xorm:"comment('排序') index INT(10)"`
	DelFlag    int       `xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	Creator    int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate time.Time `xorm:"comment('创建时间') index DATETIME"`
	Updater    int64     `xorm:"comment('更新者') BIGINT(20)"`
	UpdateDate time.Time `xorm:"comment('更新时间') DATETIME"`
}
package entity

import "time"

type SysDept struct {
	Id         int64     `json:"id" xorm:"pk comment('id') BIGINT(20)"`
	Pid        int64     `json:"pid" xorm:"comment('上级ID') index BIGINT(20)"`
	Pids       string    `json:"pids" xorm:"comment('所有上级ID，用逗号分开') VARCHAR(500)"`
	Name       string    `json:"name" xorm:"comment('部门名称') VARCHAR(50)"`
	Sort       int       `json:"sort" xorm:"comment('排序') INT(10)"`
	DelFlag    int       `json:"delFlag" xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	Creator    int64     `json:"creator" xorm:"comment('创建者') BIGINT(20)"`
	CreateDate time.Time `json:"createDate" example:"2020-05-27" xorm:"comment('创建时间') index DATETIME"`
	Updater    int64     `json:"updater" xorm:"comment('更新者') BIGINT(20)"`
	UpdateDate time.Time `json:"updateDate" example:"2020-05-27" xorm:"comment('更新时间') DATETIME"`
}

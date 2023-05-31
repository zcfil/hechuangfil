package entity

import (
	"time"
)

type SysParams struct {
	Id         int64     `xorm:"pk comment('id') BIGINT(20)"`
	ParamCode  string    `xorm:"comment('参数编码') unique VARCHAR(32)"`
	ParamValue string    `xorm:"comment('参数值') VARCHAR(2000)"`
	ParamType  int       `xorm:"default 1 comment('类型   0：系统参数   1：非系统参数') TINYINT(4)"`
	Remark     string    `xorm:"comment('备注') VARCHAR(200)"`
	DelFlag    int       `xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	Creator    int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate time.Time `xorm:"comment('创建时间') index DATETIME"`
	Updater    int64     `xorm:"comment('更新者') BIGINT(20)"`
	UpdateDate time.Time `xorm:"comment('更新时间') DATETIME"`
}

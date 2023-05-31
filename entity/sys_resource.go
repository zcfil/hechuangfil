package entity

import (
	"time"
)

type SysResource struct {
	Id             int64     `xorm:"pk comment('id') BIGINT(20)"`
	ResourceCode   string    `xorm:"comment('资源编码，如菜单ID') index VARCHAR(32)"`
	ResourceName   string    `xorm:"comment('资源名称') VARCHAR(32)"`
	ResourceUrl    string    `xorm:"comment('资源URL') VARCHAR(100)"`
	ResourceMethod string    `xorm:"comment('请求方式（如：GET、POST、PUT、DELETE）') VARCHAR(8)"`
	MenuFlag       int       `xorm:"comment('菜单标识  0：非菜单资源   1：菜单资源') TINYINT(3)"`
	AuthLevel      int       `xorm:"comment('认证等级   0：权限认证   1：登录认证    2：无需认证') TINYINT(3)"`
	Creator        int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate     time.Time `xorm:"comment('创建时间') index DATETIME"`
	Updater        int64     `xorm:"comment('更新者') BIGINT(20)"`
	UpdateDate     time.Time `xorm:"comment('更新时间') DATETIME"`
}

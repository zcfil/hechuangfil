package entity
import (
	"time"
)

type SysRoleDataScope struct {
	Id         int64     `xorm:"pk comment('id') BIGINT(20)"`
	RoleId     int64     `xorm:"comment('角色ID') index BIGINT(20)"`
	DeptId     int64     `xorm:"comment('部门ID') BIGINT(20)"`
	Creator    int64     `xorm:"comment('创建者') BIGINT(20)"`
	CreateDate time.Time `xorm:"comment('创建时间') DATETIME"`
}

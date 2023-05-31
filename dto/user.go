package dto

import (
	"hechuangfil/entity"
)

//LoginDto 登录对象
type LoginDto struct{

	Username   string    `json:"username" xorm:"comment('用户名') unique VARCHAR(50)"`
	Password   string    `json:"password" xorm:"comment('密码') VARCHAR(100)"`

}

//MerchantsDto 商户对象
type MerchantsDto struct {

	entity.SysDept
	entity.SysUser
}

package entity

type Business struct {
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	BusinessName string `xorm:"not null default '' comment('名称') unique VARCHAR(20)"`
	BusinessPwd  string `xorm:"not null comment('密码') VARCHAR(50)"`
	Solt         string `xorm:"not null comment('加密盐') VARCHAR(40)"`
	Name         string `xorm:"not null comment('用户名') VARCHAR(10)"`
	Age          int    `xorm:"not null comment('年龄') TINYINT(10)"`
	Sex          int    `xorm:"not null comment('性别') TINYINT(5)"`
	Addr         string `xorm:"not null comment('地址') VARCHAR(50)"`
}

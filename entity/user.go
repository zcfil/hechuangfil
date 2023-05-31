package entity

import "time"

type User struct {
	Id			int64     `json:"userId,string" xorm:"pk comment('id') BIGINT(20)"`
	Phone	 	string    `json:"phone" xorm:"comment('手机') unique VARCHAR(20)"`
	Username  	string    `json:"username" xorm:"comment('用户名') unique VARCHAR(50)"`
	Password   	string    `json:"password" xorm:"comment('密码') VARCHAR(100)"`
	Certification int		`json:"certification" xorm:"comment('是否实名制 0未实名，1已实名') int(11)"`
	AddressId	int64    `json:"address_id" xorm:"comment('地址ID') BIGINT(20)"`
	Delflag    	int       `json:"delflag" xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	Email      	string    `json:"email" xorm:"comment('邮箱') VARCHAR(100)"`
	CreateDate time.Time `json:"createDate" xorm:"comment('创建时间') index DATETIME"`
}
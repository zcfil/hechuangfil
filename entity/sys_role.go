package entity

type SysRole struct {
	Id         int64  `json:"id" swaggerignore:"true" xorm:"pk comment('id') BIGINT(20)"`
	Name       string `json:"name" xorm:"comment('角色名称') VARCHAR(32)"`
	Remark     string `json:"remark" xorm:"comment('备注') VARCHAR(100)"`
	DelFlag    int    `json:"delFlag" xorm:"comment('删除标识  0：未删除    1：删除') index TINYINT(4)"`
	DeptId     int64  `swaggerignore:"true" xorm:"comment('部门ID') index BIGINT(20)"`
	Creator    int64  `json:"creator" xorm:"comment('创建者') BIGINT(20)"`
	CreateDate string `json:"createDate" swaggertype:"string" example:"2020-05-27" xorm:"comment('创建时间') index DATETIME"`
	Updater    int64  `json:"updater" xorm:"comment('更新者') BIGINT(20)"`
	UpdateDate string `json:"updateDate" example:"2020-05-27" xorm:"comment('更新时间') DATETIME"`
}

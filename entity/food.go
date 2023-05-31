package entity

type Food struct {
	Id         int     `xorm:"not null pk autoincr INT(10)"`
	BusinessId int     `xorm:"not null comment('业务id') INT(10)"`
	FoodName   string  `xorm:"not null comment('食物名称') VARCHAR(10)"`
	FoodPrice  float32 `xorm:"not null comment('食物价格') FLOAT(10,2)"`
	PhotoName  string  `xorm:"not null comment('图片资源访问地址') VARCHAR(20)"`
}

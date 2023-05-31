package entity

import "time"

type Orders struct {
	OrderId           string     `gorm:"column:order_id" json:"order_id"`
	Amount    	 float64    `gorm:"column:amount" json:"amount"`               //金额
	CustomerId     	string    `gorm:"column:customer_id" json:"customer_id"`               //客户id
	Remark        string    `gorm:"column:remark" json:"remark"`            //备注
	IsDel		int			`gorm:"column:is_del" json:"is_del"`	//是否删除
	Status 		int 		`gorm:"column:status" json:"status"`
	CreateTime	time.Time			`gorm:"column:create_time" json:"create_time"`	//创建时间
	UpdateTime	time.Time			`gorm:"column:update_time" json:"update_time"`	//创建时间
	Hashrate 	int64    			`gorm:"column:hashrate" json:"hashrate"`
	ExpirationDate	time.Time			`gorm:"column:expiration_date" json:"expiration_date"`	//创建时间
}

func GetOrders(mp []map[string]string){

}
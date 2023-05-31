package entity


type UserLevel struct {
	Percentreality	float64 	 `gorm:"column:percentreality" json:"percentreality"`
	CustomerId	string `gorm:"column:customer_id" json:"customer_id"`
	Accumulative	float64 `gorm:"column:accumulative" json:"accumulative"`
	AccumulativePrice	float64 `gorm:"column:accumulative_price" json:"accumulative_price"`
	Levelvalue	float64 `gorm:"column:levelvalue" json:"levelvalue"`
	LevelvaluePrice	float64 `gorm:"column:levelvalue_price" json:"levelvalue_price"`
	Percent	float64 	`gorm:"column:percent" json:"percent"`
	Levelname	string `gorm:"column:levelname" json:"levelname"`
}
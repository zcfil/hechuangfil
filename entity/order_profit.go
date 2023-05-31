package entity

type OrderProfit struct {
	ID          	 string     `gorm:"column:id" json:"id"`
	OrderId           string     `gorm:"column:order_id" json:"order_id"`
	UserId      string     `gorm:"column:user_id" json:"user_id"`
	Profits           float64     `gorm:"column:profits" json:"profits"`
	CustomerId           string     `gorm:"column:customer_id" json:"customer_id"`
}
package entity

import "time"

type SettleLog struct {
	Id 				int64 `json:"id"`
	CustomerId 		int64  `json:"customer_id"`
	CustomerIncome 	float64	`json:"customer_income"`
	ToCustomerBalance float64	`json:"to_customer_balance"`
	ToCustomerLock 	float64	`json:"to_customer_lock"`
	CustomerLockRelease float64			`json:"customer_lock_release"`
	Time 			time.Time			`json:"time"`
}
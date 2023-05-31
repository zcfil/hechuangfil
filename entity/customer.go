package entity

import "time"

type Customer struct {
	CustomerID   	int64    	`json:"customer_id" xorm:"customer_id"`
	Name 			string		`json:"name" xorm:"name"`
	Password 		string		`json:"password" xorm:"password"`
	Phone 			string		`json:"phone" xorm:"phone"`
	UserID          int			`json:"user_id" xorm:"user_id"`
	IsDel           int8		`json:"is_del" xorm:"is_del"`
	Status          int8		`json:"status"`
	CreateTime 		time.Time	`json:"create_time" xorm:"create_time"`
	UpdateTime 		time.Time	`json:"update_time" xorm:"update_time"`
	Sex 			int8		`json:"sex"`
	Identity 		string		`json:"identity"`
	Wallet 			string		`json:"wallet"`
	Balance 		float64		`json:"balance"`
	PayPassword     string 		`json:"pay_password" xorm:"pay_password"`
	WithdrawWallet 	string		`json:"withdraw_wallet"`
	ReferrerId 		int64		`json:"referrer_id"`
	InvitationCode 	string		`json:"invitation_code"`
}


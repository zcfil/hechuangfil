package entity

type Withdraw struct {
	Id         	int64     	`xorm:"pk comment('id') BIGINT(20)"`
	UserId     	int64     	`xorm:"index BIGINT(20)"`
	UserName   	string    	`xorm:"not null comment('用户姓名') VARCHAR(10)"`
	Amount		float64		`xorm:"comment('提现金额') DOUBLE(255)"`
	CoinType	string    	`xorm:"not null comment('币种类型') VARCHAR(10)"`
	Status      int       	`xorm:"comment('状态0待审，1审核通过，2拒绝提币') TINYINT(4)"`
	Datastatus  int			`xorm:"comment('删除标识  0：未删除    1：删除') TINYINT(4)"`
	Address    	int64     	`xorm:"not null comment('钱包地址') VARCHAR(80)"`
	BatchCount    	int64     	`xorm:"not null comment('批次') INT"`
	CreateTime    	string     	`xorm:"not null comment('创建时间戳') VARCHAR(20)"`
	UpdateTime    	int64     	`xorm:"not null comment('更新时间戳') BIGINT"`
}

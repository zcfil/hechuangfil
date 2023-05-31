package entity

type SysLanguage struct {
	TableName  string `xorm:"not null pk comment('表名') VARCHAR(32)"`
	TableId    int64  `xorm:"not null pk comment('表主键') index BIGINT(20)"`
	FieldName  string `xorm:"not null pk comment('字段名') VARCHAR(32)"`
	FieldValue string `xorm:"not null comment('字段值') VARCHAR(200)"`
	Language   string `xorm:"not null pk comment('语言') VARCHAR(10)"`
}

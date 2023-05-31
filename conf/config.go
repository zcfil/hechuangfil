package conf


type Config struct {
	Repo *Repo
	Mysql *MySql
	Redis *Redis
	Dysmsapi *Dysmsapi
	Project *Project
	LogFile *LogFile
	AuthCode *AuthCode
}

type Repo struct{
	PaymentImagePath string
	ProductImagePath string
	MaxSize int64
}
type Dysmsapi struct{
	AccessKeyID		 	string
	AccessKeySecret		string
	RegisterMsgTemplate string
	LoginMsgTemplate	string
	SignName	string
}

type MySql struct {
	Name string
	Host string
	Username string
	Password string
}

type Redis struct {
	Host string
	Password string
}
type Project struct {
	Host string
	LotusUrl string
	Salesmanratio float64
}

type LogFile struct {
	FileName string
}

type AuthCode struct {
	Token string
}

const(
	MSGCODETYPE_REGISTER = "0"
	MSGCODETYPE_LOGIN = "1"
	MSGCODETYPE_WITHDRAW = "2"
)
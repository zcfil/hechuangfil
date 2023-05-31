package define

const AUTH_CODE_WIDTH = 6 // 手机验证码长度

const AUTH_CODE_VALID_TIME = 60*5 // 验证码有效时间

const AUTH_CODE_URL = "https://sms-api.upyun.com/api/messages"


// 前端请求验证码类型
const(
	CLIENT_CODE_TYPE_REGISTER = "1" // 注册验证码
	CLIENT_CODE_TYPE_LOGIN    = "2" // 登录验证码
	CLIEN_CODE_TYPE_AUTH      = "3" // 验证码
)

const (
	TEMPLATE_AUTH_CODE = 5398					// 验证码
	TEMPLATE_LOGIN_CODE = 5397					// 登录验证码
	TEMPLATE_REGISTER_CODE = 5396				// 注册验证码
)
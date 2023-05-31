package conf

import (
	"github.com/go-redis/redis/v7"
	"time"
)

//token在缓存中的有效时间
const (
	//api_key名称
	API_KEY string = "token"
	//缓存中token名称
	TOKEN_PREFIX_USER string = "user_login:prefix_token_"
	TOKEN_PREFIX_ADMIN string = "admin_login:prefix_token_"
	//token有效期一天
	TOKEN_EFFECT_TIME time.Duration = 24 * 60 * 60 *time.Second*7
	//miner key名称
	MINERS_TOKEN string = "miners_token:miner_token_"
	//miner key有效期
	MINERS_EFFECT_TIME time.Duration = time.Hour * 24
	//user roles name
	ROLES_PREFIX string = "user_token:role_token_"

	//验证码
	PREFIX_USER_LOGIN_CODE string = "login_code:prefix_token_"
	PREFIX_USER_REGISTER_CODE string = "register_code:prefix_token_"
	PREFIX_USER_WITHDRAW_CODE string = "register_code:prefix_token_"
	//验证码有效期
	CODE_EFFECT_TIME time.Duration = 60 * 5 *time.Second

	PRODUCTID_COUNT = "productid_count:prefix_token_"
)

func NewRedisClient(r *Redis) func() (*redis.Client, error) {

	return func() (*redis.Client, error) {

		client := redis.NewClient(&redis.Options{
			Addr: r.Host,
			DB: 0,
		})
		return client, nil
	}
}

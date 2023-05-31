package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SignKey = []byte("secdsdsret")
)

//LoginClaims jwt构建凭证
type LoginClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Sign 签名生成token字符串
func Sign(username string,expire int64) (string, error) {

	claims := LoginClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			//Audience:  "",
			ExpiresAt: time.Now().Unix() + expire,
			//Id:        "",
			//IssuedAt:  0,
			Issuer: "yun",
			//NotBefore: 0,
			//Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SignKey)

	return tokenString, err

}

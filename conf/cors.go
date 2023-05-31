package conf

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Cors 插件配置链接https://github.com/gin-contrib/cors
func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE","OPTIONS"},
		AllowHeaders:     []string{"token", "content-type"},
		AllowCredentials: true,
		AllowFiles:       true,
	})
}

package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
	"hechuangfil/conf"
	"hechuangfil/result"
	"hechuangfil/utils"
)


func NetMiners(client *redis.Client) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		key := ctx.Request.URL.RequestURI()
		fmt.Println(ctx.Request.URL.RequestURI())
		i, err := client.Exists(conf.MINERS_TOKEN +key).Result()
		if err != nil || i ==0{
			//pass数据库
			ctx.Next()
			return
		}
		res, err := client.Get(conf.MINERS_TOKEN + key).Result()
		if err != nil {
			fmt.Println("data from cache err:",err)
			//pass数据库
			ctx.Next()
			return
		}
		var list utils.Pagination
		err = json.Unmarshal([]byte(res), &list)
		if err != nil {
			fmt.Println("序列化数据 err:",err)
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusOK,result.Ok(list))
		ctx.Abort()
		return
	}
}


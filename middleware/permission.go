package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"log"
	"hechuangfil/conf"
	"hechuangfil/entity"
	"hechuangfil/result"
)

//CheckPermission 用户的角色控制中间件
func CheckPermission(client *redis.Client,enforcer *casbin.Enforcer) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		//sysUser, _ := utils.GetSubject(client, ctx)
		//redis中获取roles等信息
		token := ctx.Request.Header.Get("token")
		res, err := client.Get(conf.ROLES_PREFIX+token).Result()
		if err != nil {
			ctx.JSON(200,result.Fail(err))
			ctx.Abort()
			return
		}
		var roles []entity.YungoRole
		err = json.Unmarshal([]byte(res), &roles)
		if err != nil {
			ctx.JSON(200,result.Fail(err))
			ctx.Abort()
			return
		}
		for _,sub := range roles {
			fmt.Println("role:",sub)
			log.Printf("请求资源:%s,请求的动作或方法:%s",obj,act)
			//判断策略中是否存在
			if ok, _ := enforcer.Enforce(sub.Role, obj, act); ok {
				ctx.Next()
				return
			} else {
				continue
			}
		}
		ctx.JSON(200,result.Fail(errors.New("权限不通过")))
		ctx.Abort()
		return
		/*var sub string
		if sysUser.SuperAdmin == 1 {
			sub ="admin"
		}else {
			sub = "member"
		}*/
		//用户角色
		//sub := "super"
		/*log.Printf("请求资源:%s,请求的动作或方法:%s",obj,act)
		//判断策略中是否存在
		if ok, _ := enforcer.Enforce(sub, obj, act); ok {
			ctx.Next()
		} else {
			ctx.JSON(200,result.Fail(errors.New("权限不通过")))
			ctx.Abort()
			return
		}*/

	}
}


package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"hechuangfil/result"
	"net/http"
)

type News struct {
	*xorm.Engine
}


func NewNews(x *xorm.Engine) *News {
	news := new(News)
	news.Engine = x
	return news
}


func (this *News) GetNews(c *gin.Context) {
	sql := `select id, title from news where status=1`
	mapList, err := this.Engine.QueryString(sql)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(mapList))
}

func (this *News) GetNewsContent(c *gin.Context) {
	id := c.Request.FormValue("id")
	sql := `select * from news where id=%s and status=1`
	sql = fmt.Sprintf(sql, id)
	mapList, err := this.Engine.QueryString(sql)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(mapList[0]))
}
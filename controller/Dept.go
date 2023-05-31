package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"go.uber.org/dig"
	"log"
	"hechuangfil/entity"
	"hechuangfil/result"
	"strconv"
)

/**
  组织机构
 */
type Dept struct {
	dig.In
	*xorm.Engine
	*redis.Client
}

//@description 新增组织机构
//@accept json
//@Param sysDept body entity.SysDept true "组织"
//@Success 200 {object} gin.H
//@router /dept/add [post]
//@Security ApiKeyAuth
//func (dept *Dept) InsertDept(ctx *gin.Context) {
//
//	var sysDept entity.SysDept
//
//	err := ctx.ShouldBindJSON(&sysDept)
//
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, err.Error())
//		ctx.Abort()
//		return
//	}
//	sysDept.Id = utils.Node().Generate().Int64()
//	//获取当前用户的ID
//	currentUser, err := utils.GetSubject(dept.Client,ctx)
//
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, err.Error())
//		ctx.Abort()
//		return
//	}
//	sysDept.Creator = currentUser.Id
//	sysDept.Updater = currentUser.Id
//	sysDept.CreateDate = time.Now()
//	sysDept.UpdateDate = time.Now()
//	_, err = dept.Engine.Insert(&sysDept)
//	if err != nil {
//
//		ctx.JSON(http.StatusOK, gin.H{"code": -1, "data": nil, "msg": err.Error()})
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": sysDept, "msg": "ok"})
//}

//@description 查询待审核的组织
//@accept json
//@Param pageNo query int true "当前页数"
//@Param pageSize query int true "每页数据"
//@Success 200 {object} gin.H
//@router /dept/verify [get]
//@Security ApiKeyAuth
func (dept *Dept) VerifyList(ctx *gin.Context) {

	pageNo := ctx.DefaultQuery("pageNo", "1")

	pageSize := ctx.DefaultQuery("pageSize", "10")

	var list = make([]entity.UserDept, 0)
	//执行 更新操作
	//result, _ := dept.Engine.Exec("select * from sys_user")
	var sql = `SELECT 
		sys_user.id,
		sys_user.username,
		sys_user.real_name,
		sys_user.gender,
		sys_user.mobile,
		sys_user.email,
		sys_user.head_url,
		sys_user.create_date,
		sys_user.update_date,
		sys_user.dept_id,
		sys_dept.name
	FROM
		sys_user,
		sys_dept
	WHERE
		sys_user.dept_id = sys_dept.id
	AND sys_user.verified = 0
	LIMIT ?,?`

	number, _ := strconv.Atoi(pageNo)
	size, _ := strconv.Atoi(pageSize)

	dept.Engine.Table("sys_user").SQL(sql, (number-1)*size, size).Find(&list)
	//var coutSql = `SELECT
	//	count(*)
	//FROM
	//	sys_user,
	//	sys_dept
	//WHERE
	//	sys_user.dept_id = sys_dept.id
	//AND sys_user.verified = 0`

	log.Println("list:", list)

	//total, _ := dept.Engine.Table("sys_user").SQL(coutSql).Count()
	//pagination := utils.NewPagination(number, size, int(total), list)
	ctx.JSON(200, result.Ok(list))

}

//@description 审核通过
//@accept json
//@Param id path int true "商户申请人id"
//@Success 200 {object} gin.H
//@router /dept/check/{id} [put]
//@Security ApiKeyAuth
func (dept *Dept) CheckDept(ctx *gin.Context) {

	id := ctx.Param("id")

	exec, err := dept.Engine.Exec("update `sys_user` set verified = 1 where id=?", id)

	if err != nil {

		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	affected, err := exec.RowsAffected()

	if err != nil {

		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}

	if affected != 1 {
		ctx.JSON(200, result.Fail(errors.New("更新失败")))
		ctx.Abort()
		return
	}
	log.Println("受影响得行数:", affected)
	ctx.JSON(200, result.Ok(nil))
}

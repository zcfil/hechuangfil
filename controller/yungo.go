package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"go.uber.org/dig"
	"log"
	"net/http"
	"strconv"
	"time"
	"hechuangfil/conf"
	"hechuangfil/entity"
	"hechuangfil/result"
	"hechuangfil/utils"
)

type Yungo struct {
	dig.In
	*xorm.Engine
	*redis.Client
	*conf.Repo
}


//文件上传
func (y *Yungo) Upload(ctx *gin.Context) {
	// single file
	file, err := ctx.FormFile("file")
	ftype := ctx.DefaultPostForm("ftype", "5")
	ft, err := strconv.ParseInt(ftype, 10, 32)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	fmt.Println("ftype----------",ftype)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusOK, err.Error())
		return
	}
	log.Println("file-name:", file.Filename)
	// Upload the file to specific dst.
	ctx.SaveUploadedFile(file, y.ProductImagePath+file.Filename)
	//lotus client import
	if err != nil {
		ctx.JSON(200, err.Error())
		ctx.Abort()
		return
	}

	user, err := utils.GetSubject(y.Client,ctx)
	if err != nil {
		ctx.JSON(200, err.Error())
		ctx.Abort()
		return
	}
	//存到databae
	yfile := &entity.YungoFile{
		MemberId:   user.CustomerID,
		Type:       int(ft),
		Filename:   file.Filename,
		FileSize:   strconv.FormatInt(file.Size, 10),
		Price:      "0.00000001",
		MinerId:    "",
		DealId:     "",
		CreateTime: time.Now().UTC(),
		UpdateTime: time.Now().UTC(),
	}
	_, err = y.Insert(yfile)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, result.Ok(yfile))
}

//文件类型列表
func (y *Yungo) ListFileTypes(ctx *gin.Context) {

	var types  []entity.YungoFileType
	err := y.Find(&types)
	if err !=nil {
		log.Println("lotus client local:", err)
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	ctx.JSON(200,result.Ok(types))
}

func (y *Yungo) DownLoad(ctx *gin.Context) {

	fid := ctx.DefaultQuery("id","")
	if fid ==""{
		ctx.JSON(200, result.Fail(errors.New("参数不能为空")))
		ctx.Abort()
		return
	}

	user, err := utils.GetSubject(y.Client, ctx)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	var file entity.YungoFile
	has, err := y.Engine.Where("id = ? AND member_id = ?", fid,user.CustomerID).Get(&file)
	if err != nil {
		ctx.JSON(200, result.Fail(err))
		ctx.Abort()
		return
	}
	if has {
		//查询文件
		ctx.File(fmt.Sprintf("%s%s",y.Repo.ProductImagePath,file.Filename))
	}else {
		ctx.JSON(200, result.Fail(errors.New("文件不存在")))
		ctx.Abort()
		return
	}
}


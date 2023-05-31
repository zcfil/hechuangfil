package controller

import (
	"hechuangfil/service"
)

type Finance struct {
	*service.FinanceService
}
func NewFinance(s service.FinanceService) *Finance {

	return &Finance{
		&s,
	}
}

//func (fi *Finance)Upload(ctx *gin.Context){
//
//	file, err := ctx.FormFile("file")
//	if err != nil {
//
//	}
//	tokenStr := ctx.GetHeader(conf.API_KEY)
//	user, _ := utils.GetSubjectByTokenStr(tokenStr, fi.Client,conf.TOKEN_PREFIX_USER)
//	f,_ := file.Open()
//
//	if err = fi.UploadApply(f,file.Size,user);err!=nil{
//		ctx.JSON(http.StatusOK, result.Fail(err))
//	}
//	ctx.JSON(http.StatusOK, result.Ok("导入成功"))
//}

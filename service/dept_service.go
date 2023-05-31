package service

import "github.com/go-xorm/xorm"

type DeptService struct {
	*xorm.Engine
}

//NewDeptService
func NewDeptService(engine *xorm.Engine) *DeptService {

	return &DeptService{
		Engine: engine,
	}
}


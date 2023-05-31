package utils

import (
	"encoding/json"
	"strconv"
)

//Pagination 分页数据
type Pagination struct {
	PageIndex   int         `json:"pageIndex"`
	PageSize int         `json:"pageSize"`
	Data     interface{} `json:"data"`
	Total    int         `json:"total"`
}
//type PageData struct {
//	PageNo   int         `json:"pageNo"`
//	PageSize int         `json:"pageSize"`
//	List     interface{} `json:"list"`
//	Total    int         `json:"total"`
//}
type PageData struct {
	PageIndex   int         `json:"pageIndex"`
	PageSize int         `json:"pageSize"`
	Data     interface{} `json:"data"`
	Total    int         `json:"total"`
	Code	 int		`json:"code"`
	Summation     interface{} `json:"summation"`
}
func NewPagination(pageIndex, pageSize, total int, list interface{}) *Pagination {

	return &Pagination{
		PageIndex:   pageIndex,
		PageSize: pageSize,
		Data:     list,
		Total:    total,
	}
}

//Pagination
func NewPageData(param map[string]string, data interface{}) *PageData {
	no,_ := strconv.Atoi(param["pageIndex"])
	size,_ := strconv.Atoi(param["pageSize"])
	to,_ := strconv.Atoi(param["total"])
	return &PageData{
		PageIndex:   no,
		PageSize: size,
		Data:     data,
		Total:    to,
		Code: 200,
		//Summation: param["summation"],
	}
}
func NewPageDataTotal(param map[string]string, data interface{},total interface{}) *PageData {
	no,_ := strconv.Atoi(param["pageIndex"])
	size,_ := strconv.Atoi(param["pageSize"])
	to,_ := strconv.Atoi(param["total"])
	return &PageData{
		PageIndex:   no,
		PageSize: size,
		Data:     data,
		Total:    to,
		Code: 200,
		Summation: total,
	}
}
func (p *Pagination)String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}


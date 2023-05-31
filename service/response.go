package service

import (
	"encoding/json"
	"errors"
	"hechuangfil/conf"
	"hechuangfil/entity"

	//"go.uber.org/dig"
	//"hechuangfil/conf"
	"io/ioutil"
	"net/http"
)

//请求响应
type Response struct {
	*conf.Project
}

const (
	WALLET_NEW = "/LAN/walletNew"
)
func NewResponse(prf *conf.Project)*Response{
	return &Response{
		prf,
	}
}

func (res *Response) Get(url string)(*entity.Response,error){
	resp,err := http.Get(url)
	if err!=nil{
		return nil,err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	response := entity.Response{}
	// 解析json数据并将数据存储在response结构体中
	json.Unmarshal(body, &response)
	return &response,nil
}

func (res *Response) WalletNew()(string,error){
	r,err := res.Get(res.LotusUrl+ WALLET_NEW)
	if err!=nil{
		return "",err
	}
	if r.Code!=200{
		err = errors.New(r.Msg)
		return "",err
	}
	return r.Data.(string),nil
}
package entity


//请求响应
type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 错误消息
	Msg string `json:"msg"`
}
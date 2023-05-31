package entity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)


//测试 time.Time类型的解析 以及时间格式化问题
func TestTime(t *testing.T)  {

  obj := Obj{
	  From: JsonTime(time.Now()),
  }

	bytes, _ := json.Marshal(&obj)

	fmt.Println("序列化:",string(bytes))

	 o := new(Obj)
	str := `{"from":"2020-07-11 12:54:40"}`
	err := json.Unmarshal([]byte(str), o)
	if err != nil {
		fmt.Println(err)
	}
	tt := time.Time(o.From)

	fmt.Println("反序列化结果:",tt)


}
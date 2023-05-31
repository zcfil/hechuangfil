package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Get(url string)(string)  {
	// 超时时间：5秒
	//client := &http.Client{Timeout: 5 * time.Second}
	client := &http.Client{Timeout: 5 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Errorf("%v",err)
		return err.Error()
	}

	defer resp.Body.Close()
	buffer :=make([]byte,1024)
	//var buf []byte
	result := bytes.NewBuffer(nil)
	num :=0
	for {
		n, err := resp.Body.Read(buffer[0:])
		num += n
		//buf = append(buf, buffer...)
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Println("云构返回长度：",len(result.String()),strings.Contains(result.String(),`≈¥`))
	return result.String()
}
const (
	//CNYReg = `<span data-v-34fdfde4>≈¥([0-9]*\.[0-9]+)`
	CNYReg = `"legal_currency_price":([0-9]*\.[0-9]+)`
)
func GetCNYbyFIL(){
	str := Get("https://www.mytokencap.com/currency/fil/821765876")
	mc := regexp.MustCompile(CNYReg)

	submatch := mc.FindAllStringSubmatch(str,1)
	for _, m := range submatch {
		fmt.Println("云构：",m)
	}
}

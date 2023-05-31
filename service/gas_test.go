package service

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"hechuangfil/utils"
)
func getCNYbyFIL()float64{
	str := utils.Get(URL)
	mc := regexp.MustCompile(CNYReg)
	submatch := mc.FindAllStringSubmatch(str,-1)
	var fil float64
	for _, m := range submatch {
		fil,_ = strconv.ParseFloat(m[1],64)
		fmt.Println("云构：",m[1],fil)
		break
	}
	return fil
}
func TestSize(t *testing.T)  {

	fmt.Println(getCNYbyFIL())

}

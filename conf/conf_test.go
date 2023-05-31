package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strconv"
	"hechuangfil/utils"
)


const URL = "https://www.mytokencap.com/currency/fil/821765876"
const (
	CNYReg = `<span data-v-34fdfde4>≈¥([0-9]*\.[0-9]+)`
)
func GetCNYbyFIL()float64{
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
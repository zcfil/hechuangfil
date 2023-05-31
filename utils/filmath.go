package utils

import "strconv"

const (
	AttoFIL = 18
	NanoFIL = 9
)

func NanoOrAttoToFIL(fil string,filtype int)(res float64){
	//大于18or9位
	if len(fil)>filtype{
		str := fil[0:len(fil)-filtype]+"."+fil[len(fil)-filtype:]
		res,_ = strconv.ParseFloat(str,64)
		return res
	}
	//小于18or9位
	str := "0."
	for i:=0;i<filtype-len(fil);i++{
		str +="0"
	}
	str = str +fil
	res,_ = strconv.ParseFloat(str,64)
	return res
}

func NanoOrAttoToFILstr(fil string,filtype int)(res string){
	//大于18or9位
	if len(fil)>filtype{
		str := fil[0:len(fil)-filtype]+"."+fil[len(fil)-filtype:]
		res = str
		return res
	}
	//小于18or9位
	str := "0."
	for i:=0;i<filtype-len(fil);i++{
		str +="0"
	}
	str = str +fil
	return res
}
package utils

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"math/rand"
	"time"
)
const (
	LOTUS_HEIGHT = "lotus_height"	//记录检查过lotus的高度
	FIL_ADDRESS = "fil_address"		//钱包地址列表
	INVITATION_CODE = "invitation_code"		//钱包地址列表
)
const code = "ABCDEFGHIJKLMNOPQRSTUVWSYZ"
const codelen = 6
func AddFilAddress(client *redis.Client,addrcustid map[string]string )error{
	by,_ := client.Get(FIL_ADDRESS).Bytes()
	filAddress := make(map[string]string)
	if len(by)>0{
		_ = json.Unmarshal(by,&filAddress)
	}
	for k,v := range addrcustid{
		filAddress[k] = v
	}
	b,_ := json.Marshal(filAddress)
	if err := client.Set(FIL_ADDRESS,string(b),-1).Err();err!=nil{
		return err
	}

	return nil
}
//INVITATION_CODE
func AddInvitationCode(client *redis.Client,code map[string]string )error{
	by,_ := client.Get(INVITATION_CODE).Bytes()
	invitationCode := make(map[string]string)
	if len(by)>0{
		_ = json.Unmarshal(by,&invitationCode)
	}
	for k,v := range code{
		invitationCode[k] = v
	}
	b,_ := json.Marshal(invitationCode)
	if err := client.Set(INVITATION_CODE,string(b),-1).Err();err!=nil{
		return err
	}

	return nil
}
func GetInvitationCode(client *redis.Client)map[string]string{
	by,_ := client.Get(INVITATION_CODE).Bytes()
	invitationCode := make(map[string]string)
	if len(by)>0{
		_ = json.Unmarshal(by,&invitationCode)
	}

	return invitationCode
}
func SetInvitationCode(client *redis.Client,userid string)(string,error){
	rand.Seed(time.Now().UnixNano())
	icode := GetInvitationCode(client)
	cde := make(map[string]string)
	for{
		codestr := ""
		for i:=0;i<codelen;i++{
			codestr += string(code[rand.Intn(len(code))] )
		}
		if icode[codestr]==""{
			cde[codestr] = userid
			err := AddInvitationCode(client,cde)
			return codestr,err
		}
	}
}






const MaxAddressStringLength = 2 + 84
const MainnetPrefix = "f"
const TestnetPrefix = "t"

func CheckAddress(a string)bool{
	if len(a) > MaxAddressStringLength || len(a) < 3 {
		return false
	}

	if string(a[0]) != MainnetPrefix && string(a[0]) != TestnetPrefix {
		return false
	}
	switch a[1] {
	case 1:
	case 2:
	case 3:
	default:
		return false
	}
	return true
}
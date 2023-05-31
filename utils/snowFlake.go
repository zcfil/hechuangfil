package utils

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var node *snowflake.Node
var once *sync.Once = &sync.Once{}
var lock *sync.Mutex =&sync.Mutex{}

//snowflake算法根据node获取唯一ID int64位
func Node() *snowflake.Node {

	once.Do(func() {
		node, _ = snowflake.NewNode(1)
	})
	//if node == nil{
	//	lock.Lock()
	//	defer lock.Unlock()
	//	 if node == nil{
	//
	//	 	node,_= snowflake.NewNode(1)
	//	 }
	//}
	return node
}

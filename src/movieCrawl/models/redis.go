package models

import (
	"github.com/astaxie/goredis"
)

//封装一下redis的操作
//redis链接实例
var (
	client goredis.Client
)

const (
	//key
	URL_QUEUE = "url_queue"
	//访问过的url存放的set
	URL_VISITED_SET = "url_visited_set"
)

//获得链接
func ConnectRedis(addr string){
	client.Addr = addr
}

//往list中添加URL
func PutInQueue(url string){
	client.Lpush(URL_QUEUE,[]byte(url))
}

//从list中取末尾元素，
func PutOutQueue()string{
	//也可以用Brpop，可以监听多个队列，可以设置超时时间
	byte,err :=client.Rpop(URL_QUEUE)
	if err != nil {
		panic(err)
	}
	return string(byte)
}

//获取list的长度
func GetQueueLength()int{
	length,err := client.Llen(URL_QUEUE)
	if err != nil{
		panic(err)
		return 0
	}
	return length
}

//向set中添加url
func AddUrl(url string){
	client.Sadd(URL_VISITED_SET,[]byte(url))
}
 //检查url是否访问过
 func IsVisited(url string) bool{
 	//判断在set中是否存在
 	bIsVisited,err := client.Sismember(URL_VISITED_SET,[]byte(url))
	 if err != nil {
		 return false
	 }else {
		 return bIsVisited
	 }
 }

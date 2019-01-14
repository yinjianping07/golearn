package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"strconv"
)

type RedisClientController struct {
	beego.Controller
}

var productKey = "iphone8"

func Kill(client *redis.Client){
	client.Watch(func(tx *redis.Tx) error {
		return nil
	},productKey)

	value,err1 := client.Get(productKey).Result()
	checkErr(err1)
	num,err2 := strconv.Atoi(value)
	checkErr(err2)
	if num<=100&&num>0 {
		pipe := client.TxPipeline()
		defer pipe.Close()
		pipe.IncrBy(productKey,-1)
		_,err := pipe.Exec()
		checkErr(err)

		fmt.Println("抢购成功")
	}else {
		fmt.Println("抢购失败")
	}
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}



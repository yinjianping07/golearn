package models

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init(){
	options := redis.Options{
		Addr : "127.0.0.1:6379",
		Password : "",
		DB : 0,
	}
	redisClient = redis.NewClient(&options)
	fmt.Println(redisClient.Ping().Result())
}

func GetClient()(*redis.Client){
	return redisClient
}


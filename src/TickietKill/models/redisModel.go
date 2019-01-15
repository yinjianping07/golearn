package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

const(
	//low_list
	LOW_LIST = "low_queue"

	//mid_list
	MID_LIST = "mid_queue"

	//high_list
	HIGH_LIST = "high_queue"
)

//返回redis连接池
func NewRedisPool()(*redis.Pool){
	return &redis.Pool{
		MaxIdle:20,
		IdleTimeout: 240 * time.Second,
		Dial:func()(redis.Conn,error){
			c,err := redis.DialURL(beego.AppConfig.String("redis::redisUrl"))
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//密码
			if _,autErr := c.Do("AUTH",beego.AppConfig.String("redis::redisPassword"));
				autErr !=nil{
				return nil, fmt.Errorf("redis auth password error: %s", autErr)
			}
				return c,err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s",err)
			}
			return nil
		},
	}
}

func PutInLowQueue(userName string)error{
	c := NewRedisPool().Get()
	_,err := c.Do("LPUSH",LOW_LIST,userName)
	if err != nil {
		return fmt.Errorf("LPUSH error: %s",err)
	}
	return err
}
func PutInMidQueue(userName string)error{
	c := NewRedisPool().Get()
	_,err := c.Do("LPUSH",MID_LIST,userName)
	if err != nil {
		return fmt.Errorf("LPUSH error: %s",err)
	}
	return err
}
func PutInHighQueue(userName string)error{
	c := NewRedisPool().Get()
	_,err := c.Do("LPUSH",HIGH_LIST,userName)
	if err != nil {
		return fmt.Errorf("LPUSH error: %s",err)
	}
	return err
}
func PutOutLowQueue()string{
	c := NewRedisPool().Get()
	userName,err := redis.String(c.Do("RPOP",LOW_LIST))
	if err != nil {
		return ""
	}
	return userName
}
func PutOutMidQueue()string{
	c := NewRedisPool().Get()
	userName,err := redis.String(c.Do("RPOP",MID_LIST))
	if err != nil {
		return ""
	}
	return userName
}
func PutOutHgihQueue()string{
	c := NewRedisPool().Get()
	userName,err := redis.String(c.Do("RPOP",HIGH_LIST))
	if err != nil {
		return ""
	}
	return userName
}

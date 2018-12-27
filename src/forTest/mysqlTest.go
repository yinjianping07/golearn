package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
type MovieInfo struct {
	Id int64
	Movie_id int64
	Movie_name string
	Movie_picture string
	Movie_director string
	Movie_writer string
	Movie_country string
	Movie_main_character string
	Movie_type string
	Movie_on_time string
	Movie_span string
	Movie_grade string
	_Create_time string
}

func main(){

}

//连接数据库
func connect() orm.Ormer{
	orm.RegisterDataBase("default","mysql",`root:123456@tcp(192.168.117.132:3306)/go?charset=utf8`)
	//注册定义的model
	orm.RegisterModel(new(MovieInfo))
	//获得数据库连接
	if db := orm.NewOrm();db == nil {
		log.Fatal(db)
		return nil
	} else {
		return db
	}
}

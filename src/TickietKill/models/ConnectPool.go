package models

import "github.com/astaxie/beego/orm"

func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
}

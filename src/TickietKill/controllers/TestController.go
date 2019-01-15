package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get(){

	//cfg,err := config.NewConfig("ini","F:/go/gopath/src/TickietKill/conf/app.conf")

	mysql := beego.AppConfig.String("mysql::mysqlUrl")

	this.Ctx.WriteString(mysql)
}

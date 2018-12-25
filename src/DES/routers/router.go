package routers

import (
	"DES/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test",&controllers.DesController{},"get:Des")
}
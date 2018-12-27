package routers

import (
	"DES/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test",&controllers.DesController{},"get:DesEnter;post:Post")
    beego.Router("/testDES/up",&controllers.FileOptUploadController{},"get:Get;post:Post")
    beego.Router("/testDES/down",&controllers.FileOptDownloadController{},"get:Get")
}
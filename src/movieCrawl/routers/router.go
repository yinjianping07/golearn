package routers

import (
	"movieCrawl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //配置router
    beego.Router("/movieCrawl", &controllers.MovieCrawlController{},"get:MovieCrawl")
}

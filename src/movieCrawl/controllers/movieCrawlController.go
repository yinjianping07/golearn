package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"movieCrawl/models"
	"time"
)

type MovieCrawlController struct {
	beego.Controller
}

func (c *MovieCrawlController) MovieCrawl() {
	//爬虫入口
	rootUrl := "https://movie.douban.com/"

	//电影
	var movieinfo models.MovieInfo
	//连接redis
	models.ConnectRedis(beego.AppConfig.String("redisUrl"))

	models.PutInQueue(rootUrl)
	for{
		if models.GetQueueLength() ==0 {
			break //如果队列为空，结束当前循环
		}
		tempUrl := models.PutOutQueue()
		//判断是否访问过,如果访问过，直接不用往下走了
		if models.IsVisited(tempUrl) {
			continue
		}
		sMovieHtml := GetUrlHtml(tempUrl)
		if models.FindMovieName(sMovieHtml) != "" {
			//入库
			movieinfo = models.FindAMovie(sMovieHtml,models.FindMovieId(tempUrl))
			models.AddMovie(&movieinfo)
		}
		//提取该页面的所有连接
		urls := models.FindMovieURLs(sMovieHtml)
		for _,value := range urls{
			models.PutInQueue(value)
			c.Ctx.WriteString("<br>"+value+"</br>")
		}

		//放进set里面去
		models.AddUrl(tempUrl)
		time.Sleep(time.Second)//防止爬虫太快
	}
	c.Ctx.WriteString("end of crawl")
}

//封装一下从url中得到html过程
func GetUrlHtml(url string)(string){
	res := httplib.Get(url)
	UrlHtml,err := res.String()
	if err != nil {
		panic(err)
		return ""
	}else {
		return UrlHtml
	}
}

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type FileOptUploadController struct {
	beego.Controller
}

//页面
func (c *FileOptUploadController) Get() {
	c.TplName = "page"
}

//上传文件
func (this *FileOptUploadController) Post(){
	//image对应上传文件的name属性
	f,h,_ := this.GetFile("image")

	defer f.Close()
	filename := h.Filename
	arr := strings.Split(filename,":")
	if len(arr) > 1 {
		index := len(arr) - 1
		filename = arr[index]
	}
	fmt.Println("文件名称")
	fmt.Println(filename)

	//保存文件到指定的位置
	//static/uploadfile,这个是文件的地址，第一个static前面不要有/
	this.SaveToFile("image", path.Join("static/uploadfile",filename))
	//显示在本页面，不做跳转操作
	this.TplName = "page.l"
}

//下载文件
type FileOptDownloadController struct {
	beego.Controller
}
func (this *FileOptDownloadController) Get() {
	//图片,text,pdf文件全部在浏览器中显示了，并没有完全的实现下载的功能
	//this.Redirect("/static/img/1.jpg", 302)

	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	this.Ctx.Output.Download("static/img/01.jpg","ha.jpg")
}




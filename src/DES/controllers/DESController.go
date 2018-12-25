package controllers

import (
	"DES/models"
	"fmt"
	"github.com/astaxie/beego"
)

type DesController struct {
	beego.Controller
}

func (c *DesController) Des(){
	src := c.GetString("src")
	key := c.GetString("key")

	//加密
	DESString := fmt.Sprintf("%x",models.DesEnCrypt([]byte(src),[]byte(key)))
	fmt.Println(string(DESString))
	c.Ctx.WriteString(string(DESString))

}

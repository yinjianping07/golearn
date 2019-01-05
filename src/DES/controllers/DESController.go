package controllers

import (
	"DES/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type DesController struct {
	beego.Controller
}

func (c *DesController) DesEnter(){
	c.TplName = "page.html"
}

//不处理cookie
func (c *DesController) Post(){
	test := models.Test{}

	if err := c.ParseForm(&test);err != nil {
		//err process
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(test)


	if test.Status == "0" {
		temp := models.DesTable{}
		temp.Src = test.Src
		temp.Key = test.Key
		DesEnString := string(fmt.Sprintf("%x", models.DesEnCrypt([]byte(test.Src),[]byte(test.Key))))
		temp.Decrypt = DesEnString
		//入库
		models.AddDes(&temp)
		c.Data["json"] = map[string]interface{}{
			"success":0,
			"TimeStamp":time.Now().UnixNano() / 1e6,
			"Src":test.Src,
			"Key":test.Key,
			"cipherText":DesEnString,
			//"解密":DesString,
		}
		c.ServeJSON()
		return
	} else {
		DesString := models.FindDes(test.Src)

		fmt.Println(DesString.Key)
		DesTest := string(models.DesDecrypt(models.DesEnCrypt([]byte(DesString.Src),[]byte(DesString.Key)), []byte(DesString.Key)))
		c.Data["json"] = map[string]interface{}{
			"success":0,
			"TimeStamp":time.Now().UnixNano() / 1e6,
			"Src":DesString.Decrypt,
			"Key":DesString.Key,
			//"cipherText":DesEnString,
			"Decrypt":DesTest,
		}
		//fmt.Println(c.Data)
		c.ServeJSON()
		return
	}
	//data := c.Ctx.Input.RequestBody
	//fmt.Printf("%x",data)
	//json.Unmarshal(data,&test)


	//test := Test{}
	//if err := c.ParseForm(&test);err != nil {
	//	//err process
	//	fmt.Println(err.Error())
	//	panic(err)
	//}
	//加密
	//fmt.Printf("test：%c",test)
	//DesEnString := string(fmt.Sprintf("%x", models.DesEnCrypt([]byte(test.Src),[]byte(test.Key))))
	//fmt.Println(fmt.Sprintf("%x", models.DesEnCrypt([]byte(test.Src),[]byte(test.Key))))
	//DesString := string(models.DesDecrypt(models.DesEnCrypt([]byte(test.Src),[]byte(test.Key)), []byte(test.Key)))
	//c.Ctx.WriteString("加密后的密文是：	"+string(fmt.Sprintf("%x",models.DesEnCrypt([]byte(test.Src),[]byte(test.Key)))))
	//c.Data["json"] = map[string]interface{}{"success":0,
	//						                 "TimeStamp":time.Now().UnixNano() / 1e6,
	//										 "明文": test.Src,
	//					 				     "密钥": test.Key,
	//       	                             "加密后的密文是：":DesEnString,
	//                                       "解密":DesString,
	//       	                             }
	//c.ServeJSON()
	//return
	//c.Ctx.WriteString("密文："+DesEnString)
	//c.Ctx.WriteString("解密后："+DesString)
}

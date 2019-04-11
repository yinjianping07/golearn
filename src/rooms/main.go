package main

import (
	"fmt"
	"rooms/models"
	_ "rooms/routers"
	"strconv"
	"time"
)

func main() {

	//urls := models.URLs()
	//for _,value := range urls {
	//	time.Sleep(time.Second*1)
	//	fmt.Println(value)
	//	aroom := models.FindAARoom(value)
	//	models.AddARoom(&aroom)
	//}

	//先拿异常数据映射
	for id,url := range models.ExceptionUrls(){
		aroom := models.FindAARoom(url.(string))
		//将id转换int
		relId,_ := strconv.Atoi(id)
		fmt.Println(relId,"-",url.(string))
		boolean := models.UpdateMassage(relId,&aroom)
		fmt.Println(boolean)
		time.Sleep(time.Second*1)
	}

	//fmt.Println(models.FindAARoom("https://beijing.anjuke.com/prop/view/A1611874246"))


	//tempRoom := models.FindAARoom("https://guangzhou.anjuke.com/prop/view/A1609816767")
	//fmt.Println(tempRoom)
	//models.AddARoom(&tempRoom)

	//for _,value := range models.URLs(){
	//	roomUrlHtml := models.GetUrlHtml(value)
	//	fmt.Println(models.FindARoomAge(roomUrlHtml))
	//}
	//rootUrl := "https://guangzhou.anjuke.com/sale/p1/#filtersort"
	//fmt.Println(len(models.URLs()))
	//fmt.Println(models.FindARoomAge(roomUrlHtml))
	//beego.Run()
}


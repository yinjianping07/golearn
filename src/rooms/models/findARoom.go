package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func FindAARoom(url string)(AnjuKe){

	var room AnjuKe
	strHtml := GetUrlHtml(url)
	room.Room_url = url
	room.Room_age = FindARoomAge(strHtml)
	room.Room_name = FindARoomName(strHtml)
	room.Room_area = FindARoomArea(strHtml)
	room.Room_type = FindARoomType(strHtml)
	room.Room_panorama = IsOrNot(strHtml)

	return room
}

//产权
func FindARoomAge(str string) int {
	if str == "" {
		return 0
	}
	reg := regexp.MustCompile(`<div\s*class="houseInfo-content">(.)0年</div>`)
	result := reg.FindAllStringSubmatch(str,-1)

	if result == nil {
		return -1
	}

	age,_ := strconv.Atoi(result[0][1])

	//return strconv.Atoi(result[0][1])
	return age*10
}

//小区room name
func FindARoomName(str string)string{
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*href="https://beijing\.anjuke\.com/community/view/[0-9]*".*target="_blank".*_soj=propview>(.*)</a>`)
	result := reg.FindAllStringSubmatch(str,-1)

	fmt.Println(result)

	if len(result) == 0 {
		return ""
	}
	name := result[0][1]

	return string(name)
}

//面积room area
func FindARoomArea(str string)int{
	if str == "" {
		return -1
	}
	reg := regexp.MustCompile(`<div.*class="houseInfo-content">(.*)平方米</div>`)
	result := reg.FindAllStringSubmatch(str,-1)

	if len(result) == 0 {
		return -1
	}
	temp := strings.Split(result[0][1],".")
	area,_ := strconv.Atoi(temp[0])

	return area
}

// 房屋户型 room
func FindARoomType(str string)(int){
	if str == "" {
		return -1
	}
	reg := regexp.MustCompile(`<div class="houseInfo-content">
				(.)室
				(.)厅
				(.)卫
			</div>`)
	result := reg.FindAllStringSubmatch(str,-1)
	if result == nil {
		return 0
	}
	room_type_shi,_ := strconv.Atoi(result[0][1])
	room_type_ting,_ := strconv.Atoi(result[0][2])
	room_type_wei,_ := strconv.Atoi(result[0][3])
	return room_type_shi*100+room_type_ting*10+room_type_wei
}
//生成url
func URLs() []string{
	url := "https://beijing.anjuke.com/sale/p"
	var urls []string

	for i := 0 ;i<50 ; i++  {
		temp := url + strconv.Itoa(i+1) +"/#filtersort"
		tempUrlHtml := GetUrlHtml(temp)

		for _,value := range FindAURLs(tempUrlHtml){
			urls = append(urls,value)
		}
		time.Sleep(time.Millisecond*100)
		//time.Sleep(time.Second*1)
	}
	return urls
}

func FindAURLs(str string)[]string{
	if str == "" {
		return nil
	}
	reg := regexp.MustCompile(`<a\s*data-from=""\s*data-company=""\s*title=".*"\s*href="(.*)\?.*"\s*target='_blank'\s*class="houseListTitle.">`)
	result := reg.FindAllStringSubmatch(str,-1)
	if result == nil {
		return nil
	}

	var strs []string
	for _,value := range result{
		strs = append(strs,value[1])
	}
	//fmt.Println(len(strs))
	//fmt.Println(strs)
	return strs
}

func IsOrNot(htmlStr string)int{
	reg := regexp.MustCompile(`(全景看房)`)
	result := reg.FindAllStringSubmatch(htmlStr,-1)
	//fmt.Println(result)

	if result != nil {
		return 1
	} else{
		return 0
	}
}

//封装一下从url中得到html过程
func GetUrlHtml(url string)(string){
	var correntCount,failCount = 0,0
	res := httplib.Get(url).SetTimeout(3*time.Second,3*time.Second)
	UrlHtml,err := res.String()
	if err != nil {
		//panic(err)
		failCount++
		return ""
	}else {
		correntCount++
		return UrlHtml
	}
}
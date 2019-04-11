package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var(
	db orm.Ormer
)
type BeiKe struct {
	Id int
	Room_name string
	Room_area int
	Room_type int
	Room_age int
	Room_panorama int
}

type AnjuKe struct {
	Id int
	Room_name string
	Room_area int
	Room_type int
	Room_age int
	Room_url string
	Room_panorama int
}

func init(){
	//orm.Debug = true //开启调试模式
	orm.RegisterDataBase("default","mysql","root:1234@tcp(localhost:3306)/go?charset=utf8")
	orm.RegisterModel(new(BeiKe))
	orm.RegisterModel(new(AnjuKe))
	db = orm.NewOrm()//获得数据库连接
}

func FindBRoom(id int)BeiKe{
	room := new(BeiKe)
	room.Id = id
	err := db.Read(room)
	if err != nil {
		beego.Info("查询失败", err)
	}
	return *room
}

func FindARoom(id int)AnjuKe{
	room := new(AnjuKe)
	room.Id = id
	err := db.Read(room)
	if err != nil {
		beego.Info("查询失败", err)
	}
	return *room
}

//添加安居客
func AddARoom(aroom *AnjuKe)(int64,error){
	id,err := db.Insert(aroom)
	return id,err
}

//添加贝壳
func AddBRoom(broom *BeiKe)(int64,error){
	id,err := db.Insert(broom)
	return id,err
}

//更新错误数据
func UpdateMassage(id int,aroom *AnjuKe)bool{
	res, _ := db.Raw("update anju_ke set room_name=?,room_area=?,room_type=?,room_panorama=? where id=?").Prepare()
	_,err := res.Exec(aroom.Room_name,aroom.Room_area,aroom.Room_type,aroom.Room_panorama,id)
	if err != nil {
		log.Fatal(err)
		return false
	}else {
		return true
	}
}

//处理异常数据--查询出异常数据的url映射
func ExceptionUrls()orm.Params{

	//创建map
	result := make(orm.Params)

	//查询结果映射到result中列名必须一致
	res,err := db.Raw("select id,room_url from anju_ke where room_name=?","").
		RowsToMap(&result,"id","room_url")
	if err != nil {
		panic(err)
	}
	//输出总数
	fmt.Println(res)
	for key,value := range result{
		fmt.Println(key,"---",value)
	}
	return result
}


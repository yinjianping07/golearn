package models

//负责插入数据库
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局的数据库连接资源
var(
	db_movie_info orm.Ormer
)

//对应电影表
type MovieInfo struct {
	Id int64
	Movie_id int64
	Movie_name string
	Movie_picture string
	Movie_director string
	Movie_writer string
	Movie_country string
	Movie_main_character string
	Movie_type string
	Movie_on_time string
	Movie_span string
	Movie_grade string
	_Create_time string
}
func init(){
	//orm.Debug = true //开启调试模式
	orm.RegisterDataBase("default","mysql","root:1234@tcp(localhost:3306)/go?charset=utf8")
	orm.RegisterModel(new(MovieInfo))
	db_movie_info = orm.NewOrm()//获得数据库连接
}

func AddMovie(movie_info *MovieInfo)(int64,error){
	id,err := db_movie_info.Insert(movie_info)
	return id,err
}

func SelectAllMovie()(){
	movies := new([]MovieInfo)
	_,err := db_movie_info.QueryTable("movie_comment").All(movies)
	if err != nil {
		panic(err)
	}else {
		for _,value := range *movies{
			fmt.Println(value)
		}
	}
}

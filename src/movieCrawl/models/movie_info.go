package models

//负责插入数据库
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局的数据库连接资源
var(
	db orm.Ormer
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
	orm.RegisterDataBase("default","mysql",beego.AppConfig.String("connectUrl"))
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()//获得数据库连接
}

func AddMovie(movie_info *MovieInfo)(int64,error){
	id,err := db.Insert(movie_info)
	return id,err
}

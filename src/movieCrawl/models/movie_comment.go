package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var(
	db_movie_comment orm.Ormer
)
type MovieComment struct {
	Id int64
	Movie_name string //电影名称
	User_name string //评论者昵称
	Movie_type string //电影类型
	Movie_star string //评论者打分
	Comment_time string //评论时间
	Movie_grade string //电影评分
	Good_comment string //好评率
	Middle_comment string //一般率
	Bad_comment string //差评率
	Short_comment string //短评
	_Create_time string //创建时间
}

func init(){
	orm.RegisterDataBase("default","mysql","root:1234@tcp(localhost:3306)/go?charset=utf8")
	orm.RegisterModel(new(MovieComment))
	db_movie_comment = orm.NewOrm()
}

func AddAComment(movie_comment *MovieComment) (int64,error){
	return db_movie_comment.Insert(movie_comment)
}

func SelectAll()(){
	result := db_movie_comment.Raw("select * from movie_comment;")
	fmt.Println(result)
}





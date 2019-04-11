package models

import "github.com/astaxie/beego/orm"

//余票
type Ticket struct {
		Id int `json:"id"`
		Count int `json:"count"`
}

func init(){
	orm.RegisterModel(new(Ticket))
}

func FindTicket(o orm.Ormer,id int)(int){
	ticket := new(Ticket)
	o.QueryTable("tickiet").Filter("id",id).One(&ticket,"count")
	return ticket.Count
}

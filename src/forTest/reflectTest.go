package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func(this *User)Hello(){
	fmt.Println("hello User")
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v =%v\n", f.Name, f.Type, val)

		for i := 0; i < t.NumMethod(); i++ { //这里同样通过t.NumMethod来获取它拥有的方法的数量，来决定循环的次数
			m := t.Method(i)
			fmt.Printf("%6s:%v\n", m.Name, m.Type)
		}
	}
}

func main() {
	u := User{1, "Jack", 23}
	info(u)
}
package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

// 查看匿名字段
func main() {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)

	fmt.Println(t)

	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))

	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

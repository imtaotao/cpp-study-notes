package main

import "fmt"

type person struct {
	name string
	age  int
}

func init() {
	pointerStruct()
	initStructByLocation()
}

// 在定义一些临时数据结构等场景下可以使用匿名结构体
func main() {
	var user struct {
		name string
		age  int
	}

	user.name = "chen"
	user.age = 22

	fmt.Printf("%#v\n", user)
}

// 创建指针类型结构体
func pointerStruct() {
	var taotao = new(person)

	taotao.name = "chentao"
	taotao.age = 26

	fmt.Printf("%T\n", taotao) // *main.person (是一个指针)
	fmt.Printf("%#v\n", taotao)
}

// 取结构体的地址实例化
func initStructByLocation() {
	p3 := &person{}
	fmt.Printf("initStructByLocation: %T\n", p3)
}

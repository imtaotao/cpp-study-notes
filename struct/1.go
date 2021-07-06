package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person // 实例化结构体
	p1.name = "chentao"
	p1.city = "深圳"
	p1.age = 26

	fmt.Printf("%#v\n", p1)

	keyValueInit()
}

// 使用键值对初始化结构体
func keyValueInit() {
	p := &person{
		"chentao",
		"shenzhen",
		18,
	}
	fmt.Printf("%#v\n", p)
}

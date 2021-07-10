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
	extend()
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

func extend() {
	type Animal struct {
		name string
	}

	type Dog struct {
		Feet   int8
		name   string
		Animal Animal // 通过嵌套匿名结构体实现继承
		// Animal // 省略类型是简写，属于匿名结构体，里面的属性会被扩展出来
	}

	d := &Dog{
		Feet: 4,
		name: "a name",
		Animal: Animal{ // 注意嵌套的是结构体指针
			name: "b name",
		},
	}

	fmt.Printf("%#v\n", d)
	fmt.Printf(d.name)
}

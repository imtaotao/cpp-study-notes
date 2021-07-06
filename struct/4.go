package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

// Go 语言的结构体没有构造函数，我们可以自己实现
func main() {
	p := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p)
}

// 其实就是返回一个结构体而已
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

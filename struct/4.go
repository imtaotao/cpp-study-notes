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

// 实现了一个 person 的构造函数。
// 因为 struct 是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

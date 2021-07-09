package main

import "fmt"

// 全局变量
var a, b = 1, 2

// 批量声明变量
var (
	c string
	d int
	e bool
	f float32
)

func main() {
	println(a, b)
	test()
}

func test() {
	// 简写，只能在函数内部作用域，是块级作用域
	a, b := 2, 3
	println(a, b)

	// _多用于占位，表示忽略值
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

func foo() (int, string) {
	return 10, "Q1mi"
}

// 常量只是把 var 改成 const，定义的时候必须赋值

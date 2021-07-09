package main

import "fmt"

// 空接口：
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
func main() {
	// 定义一个空接口x
	var x interface{}

	x = "string"
	fmt.Printf("type:%T, value:%v\n", x, x)

	x = 100
	fmt.Printf("type:%T, value:%v\n", x, x)

	x = true
	fmt.Printf("type:%T, value:%v\n", x, x)

	useOne(1)
	useTwo()
}

// 空接口的应用
// 1. 空接口作为函数参数，使用空接口实现可以接收任意类型的函数参数（类似 ts 里面的 any）
func useOne(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 2. 空接口作为 map 的值，使用空接口实现可以保存任意值的字典（类似 js 的对象）
func useTwo() {
	studentInfo := make(map[string]interface{})
	studentInfo["name"] = "chentao"
	studentInfo["age"] = 22
	studentInfo["married"] = false

	fmt.Println(studentInfo)
}

// Why ???
// func init() {
// 	user := map[string]interface{
// 		"name": "chen",
// 	}

// 	fmt.Println(user)
// }

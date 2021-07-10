package main

import "fmt"

// fmt 包下有 fmt.Scan、fmt.Scanf、fmt.Scanln 三个函数，可以在程序运行过程中从标准输入获取用户的输入

// func Scan(a ...interface{}) (n int, err error)
//
// 1. Scan 从标准输入扫描文本，读取由 空白符分隔 的值保存到传递给本函数的参数中，换行符视为空白符。
// 2. 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

func main() {
	var (
		name    string
		age     int
		married bool
	)

	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

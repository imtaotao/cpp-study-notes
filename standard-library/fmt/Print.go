package main

import "fmt"

// Print 系列函数会将内容输出到系统的标准输出
// 区别在于 Print 函数直接输出内容，Printf 函数支持格式化输出字符串，Println 函数会在输出内容的结尾添加一个换行符。
//
// func Print(a ...interface{}) (n int, err error)
// func Printf(format string, a ...interface{}) (n int, err error)
// func Println(a ...interface{}) (n int, err error)

func main() {
	fmt.Print("在终端打印该信息。")

	name := "chentao"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

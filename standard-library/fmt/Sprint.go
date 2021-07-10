package main

import "fmt"

// Sprint 系列函数会把传入的数据生成并返回一个字符串
//
// func Sprint(a ...interface{}) string
// func Sprintf(format string, a ...interface{}) string
// func Sprintln(a ...interface{}) string

func main() {
	s1 := fmt.Sprint("chentao") // 用变量接收返回的字符串

	name := "chentao"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("chentao")

	fmt.Println(s1, s2, s3)
}

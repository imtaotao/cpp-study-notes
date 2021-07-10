package main

import "fmt"

// Errorf 函数根据 format 参数生成格式化字符串并返回一个包含该字符串的错误
//
// func Errorf(format string, a ...interface{}) error

func main() {
	// 通常使用这种方式来自定义错误类型，例如：
	err := fmt.Errorf("这是一个错误")

	fmt.Println(err)
}

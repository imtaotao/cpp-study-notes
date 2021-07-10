package main

import "fmt"

// func Scanf(format string, a ...interface{}) (n int, err error)
//
// 1. Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
// 2. 本函数返回成功扫描的数据个数和遇到的任何错误。

func main() {
	var (
		name    string
		age     int
		married bool
	)

	// `1:chentao 2:23 3:false`，得这样输入，只有按照格式输入才行
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

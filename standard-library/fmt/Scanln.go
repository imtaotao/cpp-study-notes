package main

import "fmt"

// func Scanln(a ...interface{}) (n int, err error)
//
// 1. Scanln 类似 Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
// 2. 本函数返回成功扫描的数据个数和遇到的任何错误。

func main() {
	var (
		name    string
		age     int
		married bool
	)

	// fmt.Scanln 遇到回车就结束扫描了，这个比较常用。
	// 换行就终止，Scan 必须输入对应的个数才行
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

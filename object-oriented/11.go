package main

import "fmt"

// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：`x.(T)`
//
// x：表示类型为 interface{} 的变量
// T：表示断言 x 可能是的类型
//
// 该语法返回两个参数：
// 1. 第一个参数是 x 转化为 T 类型后的变量
// 2. 第二个值是一个布尔值，若为 true 则表示断言成功，为 false 则表示断言失败

func main() {
	var x interface{}

	x = "www.baidu.com"
	v, ok := x.(string)

	if ok {
		fmt.Println(v) // www.baidu.com
	} else {
		fmt.Println("类型断言失败")
	}
}

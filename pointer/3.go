package main

import "fmt"

// 空指针
func main() {
	var p *string // 声明一个变量指针 p

	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)

	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
}

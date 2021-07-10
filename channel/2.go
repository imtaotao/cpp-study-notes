package main

import "fmt"

// 声明的通道后需要使用 make 函数初始化之后才能使用
// 格式：`make(chan 元素类型, [缓冲大小])`

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	ch3 := make(chan []int)

	fmt.Println(ch1)
	fmt.Println(ch2)
	fmt.Println(ch3)
}

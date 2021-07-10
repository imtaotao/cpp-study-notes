package main

import "fmt"

// channel 是一种类型，一种引用类型。声明通道类型的格式如下：
// `var 变量 chan 元素类型`

var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递 int 切片的通道

func main() {
	// 通道是引用类型，通道类型的空值是 nil。
	fmt.Println(ch1) // <nil>
	fmt.Println(ch2) // <nil>
	fmt.Println(ch3) // <nil>
}

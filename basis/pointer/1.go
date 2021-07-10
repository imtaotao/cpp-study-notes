package main

import "fmt"

// & 取地址，* 根据地址取值

// `ptr := &v`  v 的类型为 T
// v: 代表被取地址的变量，类型为 T
// ptr: 用于接收地址的变量，ptr 的类型就为 *T，称做 T 的指针类型。* 代表指针。
func main() {
	a := 10
	b := &a

	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b, *b)
}

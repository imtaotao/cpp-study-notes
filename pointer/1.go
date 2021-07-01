package main

import "fmt"

// & 取地址，* 根据地址取值
func main() {
	a := 10
	b := &a

	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b, *b)
}

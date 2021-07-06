package main

import "fmt"

// 结构体的布局
func main() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}

	n := test{
		1, 2, 3, 4,
	}

	fmt.Printf("n.a %p\n", &n.a) // n.a 0xc0000ae004
	fmt.Printf("n.b %p\n", &n.b) // n.b 0xc0000ae005
	fmt.Printf("n.c %p\n", &n.c) // n.c 0xc0000ae006
	fmt.Printf("n.d %p\n", &n.d) // n.d 0xc0000ae007
}

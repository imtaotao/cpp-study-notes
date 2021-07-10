package main

import "fmt"

// %p	表示为十六进制，并加上前导的 0x

func main() {
	a := 18

	fmt.Printf("%p\n", &a)  // 0xc0000ae008
	fmt.Printf("%#p\n", &a) // c0000ae008
}

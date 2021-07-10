package main

import "fmt"

// 修改字符串，必须先转换为 byte 类型或者 rune 类型
func main() {
	s1 := "hello"

	// byteS1 := []byte(s1)
	byteS1 := []rune(s1)
	byteS1[0] = 'H'

	fmt.Println(string(byteS1))
}

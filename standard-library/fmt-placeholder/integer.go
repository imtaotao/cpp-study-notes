package main

import "fmt"

// 布尔型（就一个，不单独列出来了）
// %t	true 或 false

// %b	表示为二进制
// %c	该值对应的 unicode 码值
// %d	表示为十进制
// %o	表示为八进制
// %x	表示为十六进制，使用a-f
// %X	表示为十六进制，使用A-F
// %U	表示为 Unicode 格式：U+1234，等价于 "U+%04X"
// %q	该值对应的单引号括起来的 go 语法字符字面值，必要时会采用安全的转义表示

func main() {
	n := 65

	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
}

package main

import "fmt"

// 当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。
// rune 类型实际是一个 int32。
// Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾

func main() {
	s := "chentao"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}

	fmt.Println()

	for _, r := range s { // rune
		fmt.Printf("%v(%c) ", r, r)
	}

	fmt.Println()
}

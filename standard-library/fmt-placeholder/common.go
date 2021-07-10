package main

import "fmt"

// %v  - %v 值的默认格式标识
// %+v - 类似 %v，但是输出结构体时会添加字段名
// %#v - 值的 Go 语法表示
// %T  - 打印值的类型
// %%  - 百分号

func main() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)

	o := struct{ name string }{"chentao"}

	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)

	fmt.Printf("100%%\n")
}

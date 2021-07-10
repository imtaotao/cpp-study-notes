package main

import "fmt"

// 在 Go 语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
// 而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
// 要分配内存，就需要用到 new 和 make 函数
func main() {
	// error()
	newTest()
}

// 使用 new 函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
func newTest() {
	a := new(int) // 默认 0
	b := new(bool)

	fmt.Println(a)      // 这是一个引用类型，这里打印出来的是一个地址
	fmt.Println(*a, *b) // 0, false
}

func error() {
	var a *int
	*a = 100 // 引用没有提前分配内存，这里赋值报错
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100 // 同上
	fmt.Println(b)
}

func fixError() {
	var a *int
	a = new(int) // 使用 new 分配内存
	*a = 100
	fmt.Println(*a)
}

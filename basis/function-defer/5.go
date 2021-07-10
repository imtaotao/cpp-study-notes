package main

import "fmt"

// defer nil 函数
func test() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

// 名为 test 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。
// 然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用。
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	test()
}

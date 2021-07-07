package main

import "fmt"

func main() {
	errorTest()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "err") // send on closed channel err
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)

	// 向已关闭的通道发送数据会引发 panic
	ch <- 1
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获
func errorTest() {
	defer func() {
		println(3)
		fmt.Println(recover()) // defer panic
	}()

	defer func() {
		println(2) // defer 的函数不会被错误给终止掉
		panic("defer panic")
	}()

	println(1)
	panic("test panic")
}

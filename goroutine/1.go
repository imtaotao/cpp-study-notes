package main

import (
	"fmt"
	// "time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}

// 这一次的执行结果只打印了 `main goroutine done!`，并没有打印 `Hello Goroutine!`。为什么呢？
// 在程序启动时，Go 程序就会为 main() 函数创建一个默认的 goroutine。
// 当 main() 函数返回的时候该 goroutine 就结束了，所有在 main()函数中启动的 goroutine 会一同结束，main 函数所在的 goroutine 就像是权利的游戏中的夜王，
// 其他的 goroutine 都是异鬼，夜王一死它转化的那些异鬼也就全部 GG 了。
//
// 所以我们要想办法让main函数等一等 hello 函数，最简单粗暴的方式就是 time.Sleep 了。
func main() {
	go hello() // 启动另外一个 goroutine 去执行 hello 函数，打印会在后面，想当是异步的
	fmt.Println("main goroutine done!")
	// time.Sleep(time.Second)
}

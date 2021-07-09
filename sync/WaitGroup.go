package main

import (
	"fmt"
	"sync"
)

// 在代码中生硬的使用 time.Sleep 肯定是不合适的，Go 语言中可以使用 sync.WaitGroup 来实现并发任务的同步。 sync.WaitGroup 有以下几个方法：
// 							方法名                       功能
// 1. `(wg * WaitGroup) Add(delta int)`: 计数器 +delta
// 2. `(wg *WaitGroup) Done()`         : 计数器 -1
// 3. `(wg *WaitGroup) Wait()`         : 阻塞直到计数器变为 0

// sync.WaitGroup 内部维护着一个计数器，计数器的值可以增加和减少。
// 例如当我们启动了 N 个并发任务时，就将计数器值增加 N。
// 每个任务完成时通过调用 Done() 方法将计数器减 1。
// 通过调用 Wait() 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。

var wg sync.WaitGroup

func hello() {

	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}

func main() {
	wg.Add(1)
	go hello() // 启动另外一个 goroutine 去执行 hello 函数
	fmt.Println("main goroutine done!")
	wg.Wait()
	println("完成")
}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

// Go 语言中的操作系统线程和 goroutine 的关系：
// 1.一个操作系统线程对应用户态多个 goroutine。
// 2.go 程序可以同时使用多个操作系统线程。
// 3.goroutine 和 OS 线程是多对多的关系，即 m:n。

func main() {
	// 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	// 将逻辑核心数设为 2，此时两个任务并行执行，代码如下。

	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

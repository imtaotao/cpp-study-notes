package main

import "fmt"

// 可以通过内置的 close() 函数关闭 channel（如果不往管道里存值或者取值的时候一定记得关闭管道）
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		// 不停的从 c 中取数据
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main 结束")
}

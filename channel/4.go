package main

import "fmt"

// 这段代码能够通过编译，但是执行的时候会出现错误：`fatal error: all goroutines are asleep - deadlock!`
//
// func main() {
// 	ch := make(chan int)
// 	ch <- 10
// 	fmt.Println("发送成功")
// }
//
// 因为我们使用 `ch := make(chan int)` 创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
// 就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。
// 上面的代码会阻塞在 `ch <- 10` 这一行代码形成死锁，那如何解决这个问题呢？

// 一种方法是启用一个goroutine去接收值，例如：
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)
	go recv(ch) // 启用 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

// 无缓冲通道上的发送操作会阻塞，直到另一个 goroutine 在该通道上执行接收操作，这时值才能发送成功，两个 goroutine 将继续执行。
// 相反，如果接收操作先执行，接收方的 goroutine 将阻塞，直到另一个 goroutine 在该通道上发送一个值。
// 使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。

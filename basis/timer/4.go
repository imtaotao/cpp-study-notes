package main

import (
	"fmt"
	"time"
)

// 停止定时器
func main() {
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("定时器执行了")
	}()

	// b := timer.Stop()
	// if b {
	// 	fmt.Println("timer 已经关闭")
	// }
}

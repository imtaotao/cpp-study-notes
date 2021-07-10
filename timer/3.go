package main

import (
	"fmt"
	"time"
)

// timer 实现延时功能
func main() {
	// 1
	time.Sleep(time.Second)

	// 2
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("2秒到1")

	// 3
	<-time.After(2 * time.Second)
	fmt.Println("2秒到2")
}

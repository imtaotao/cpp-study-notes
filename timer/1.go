package main

import (
	"fmt"
	"time"
)

// Go 语言的定时器实质是单向通道
// timer 基本使用
func main() {
	timer := time.NewTimer(2 * time.Second)

	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)

	t2 := <-timer.C
	fmt.Printf("t2:%v\n", t2)
}

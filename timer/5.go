package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	// 重置定时器
	timer.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer.C)
}

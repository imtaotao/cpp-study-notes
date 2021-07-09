package main

import (
	"fmt"
	"time"
)

// 验证 timer 只能响应 1 次
func main() {
	timer := time.NewTimer(time.Second)
	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

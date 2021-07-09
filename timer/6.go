package main

import (
	"fmt"
	"time"
)

// Ticker：时间到了，多次执行
// timer 和 ticker 就想 js 里面的 setTimeout 和 setInterval
func main() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0

	// 子协程
	go func() {
		for {
			// <-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()

	// 不让子协程关闭
	for {
	}
}

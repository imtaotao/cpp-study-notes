package main

import (
	"fmt"
	"runtime"
)

// 退出当前协程(一边烧烤一边相亲，突然发现相亲对象太丑影响烧烤，果断让她滚蛋，然后也就没有然后了)
func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			// 结束协程
			runtime.Goexit()

			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
	}
}

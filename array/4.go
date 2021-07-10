package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 求数组所有元素之和
func sum(arr [10]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	// 若想做一个真正的随机数，要种子
	// seed() 种子默认是 1
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := range b {
		b[i] = rand.Intn(1000)
	}
	sum := sum(b)
	fmt.Printf("sum=%d\n", sum)
}

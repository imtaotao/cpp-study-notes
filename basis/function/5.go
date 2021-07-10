package main

import "fmt"

// 递归
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	i := 7
	fmt.Println(i, factorial(i))
}

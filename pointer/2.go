package main

import "fmt"

// 指针传递
func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}

func modify1(a int) {
	a = 100
}

func modify2(a *int) {
	*a = 100
}

package main

import "fmt"

// 返回 2 个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 也可以通过定义变量然后返回
	return func(i int) int {
			base += i
			return base
		},
		func(i int) int {
			base -= i
			return base
		}
}

func main() {
	f1, f2 := test01(10)

	// base 一直是没有消
	fmt.Println(f1(1), f2(2))

	// 此时 base 是9
	fmt.Println(f1(3), f2(4))
}

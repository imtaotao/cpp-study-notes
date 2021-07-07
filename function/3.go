package main

import "fmt"

// 外部引用函数参数局部变量
func add(base *int) func(int) int {
	return func(i int) int {
		*base += i
		return *base
	}
}

func main() {
	n := 10
	n2 := 100

	tmp1 := add(&n)
	fmt.Println(tmp1(1), tmp1(2), n)

	// 此时 tmp1 和 tmp2 不是一个实体了
	tmp2 := add(&n2)
	fmt.Println(tmp2(1), tmp2(2), n2)
}

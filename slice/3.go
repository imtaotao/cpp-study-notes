package main

import "fmt"

func test(x interface{}) {
	fmt.Printf("x: %p\n", &x)
	v, ok := x.([2]int)
	if ok {
		println("是数组")
		v[1] = 100
	} else {
		println("不是数组")
	}
}

func main() {
	a := [...]int{1, 1} // 打印的是数组，证明带上 `...` 也算带上 len，是数组
	fmt.Printf("a: %p\n", &a)

	test(a)
	fmt.Println(a) // 打印 [1, 1]，证明是值拷贝
}

package main

import "fmt"

// func test(x [2]int) { // 数组
func test(x []int) { // 切片
	fmt.Printf("x: %p\n", &x)
	x[1] = 100
}

// 切片和数组的对比
// 可以声明一个未指定大小的数组来定义切片（也就是说没有设置 len，就是切片）
func main() {
	// a := [2]int{1, 1} // 数组
	a := []int{1, 1} // 切片
	fmt.Printf("a: %p\n", &a)

	test(a)
	fmt.Println(a) // 如果是切片，打印 [1, 100]，如果是数组打印 [1, 1]，因为数组是值拷贝过去的
}

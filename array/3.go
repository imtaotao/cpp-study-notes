package main

import "fmt"

// go 语言中的函数传参都是值拷贝
// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
// 通过数组指针接收
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 数组拷贝和传参
func main() {
	var arr1 [5]int
	// 使用地址
	printArr(&arr1)
	fmt.Println(arr1)

	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

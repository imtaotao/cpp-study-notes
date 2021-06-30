package main

import "fmt"

// slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段
// 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
// 2. 切片的长度可以改变，因此，切片是一个可变的数组。
// 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
// 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
// 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
// 6. 如果 slice == nil，那么 len、cap 结果都等于 0。

var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	slice0 := arr[2:5]

	slice0[0] = 111

	slice1 := arr[2:]
	slice2 := arr[:3]

	slice2[0] = 1212

	fmt.Println(slice0, slice1, slice2)
	fmt.Println(arr)
	fmt.Println("--------")

	makeSlice()
	arr2()
}

func makeSlice() {
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
}

func arr2() {
	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100
	// s[2] += 100
	// fmt.Println(s)
}

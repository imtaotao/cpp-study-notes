package main

import "fmt"

// 1. 数组：是同一种数据类型的固定长度的序列。
// 2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
// 3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
// 4. 数组可以通过下标进行访问，下标是从 0 开始，最后一个元素下标是：len-1
//		for i := 0; i < len(a); i++ {
//
// 		}
// 		for index, v := range a {
//
// 		}
//
// 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会 panic
// 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
// 7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
// 8.指针数组 [n]*T，数组指针 *[n]T。

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "chen", 4: "tao"}

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4, 5}
	c := [5]int{2: 100, 4: 200}

	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c)

	fmt.Println("------")
	multArrays()
}

// 多维数组
var arr00 [5][3]int
var arr11 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func multArrays() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第二维度不能用 "..."

	fmt.Println(arr00, arr11)
	fmt.Println(a, b)
	fmt.Println(len(a[1]), cap(a[1])) // len 和 cap 有啥区别？？（后面的理解，len 是数组的长度，cap 是容量）
}

package main

import "fmt"

// 延迟调用
// 1. 关键字 defer 用于注册延迟调用。
// 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
// 3. 多个 defer 语句，按先进后出的方式执行（只在当前函数内）。
// 4. defer 语句中的变量，在 defer 声明时就决定了。
func main() {
	whatever := [...]struct{ a int }{{a: 0}, {a: 1}, {a: 2}, {a: 3}, {a: 4}}

	for i, v := range whatever {
		defer fmt.Println(i, v)
	}

	closerAndDefer()
}

func closerAndDefer() {
	var whatever [5]struct{}
	for i := range whatever {
		// 这里和 js 闭包一样，当打印 i 的时候，i 已经是最后一个数值了
		defer func() { fmt.Println(i) }() // 4
	}
}

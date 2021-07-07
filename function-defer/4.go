package main

import "fmt"

// 在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。
// 所以 defer closure 输出结果为 2 而不是 1
func foo() (i int) {
	i = 0
	defer func() {
		fmt.Println(i) // 2
	}()
	return 2
}

// defer 与 return
func main() {
	foo()
}

package main

// Go 编程语言中 if 语句的语法如下：
// 1. 可省略条件表达式括号。
// 2. 持初始化语句，可定义代码块局部变量。
// 3. 代码块左 括号必须在条件表达式尾部。

// if 布尔表达式 {
//   ...
// }

// if 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则不执行。

func main() {
	x := 0

	if n := "abc"; x > 0 { // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
		println(n[2])
	} else if x < 0 { // 注意 else if 和 else 左大括号位置。
		println(n[1])
	} else {
		println(n[0])
	}
}

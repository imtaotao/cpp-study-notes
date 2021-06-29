package main

// iota 是 go 语言的常量计数器，只能在常量的表达式中使用。
// iota 在 const 关键字出现时将被重置为 0。
// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
// 使用iota能简化定义，在定义枚举时很有用
func main() {
	const (
		n1 = iota
		n2
		_
		n4
	)

	println(n1, n2, n4)

	test()
}

func test() {
	const (
		a, b = iota + 1, iota + 2 // 1,2
		c, d                      // 2,3
		e, f                      // 3,4
	)

	println(a, b, c, d, e, f)
}

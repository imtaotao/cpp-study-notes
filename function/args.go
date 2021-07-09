package main

func add(args ...int) int { // 0 个或多个参数
	return 1
}

// args 是一个 slice，我们可以通过 arg[index] 依次访问所有参数,通过 len(arg) 来判断传递参数的个数
func add1(a int, args ...int) int { // 1 个或多个参数
	return 1
}

func add2(a int, b int, args ...int) int { // 2 个或多个参数
	return 1
}

// 用 interface{} 传递任意类型数据是 Go 语言的惯例用法，而且 interface{} 是类型安全的。
func add3(args ...interface{}) int {
	return 1
}

// https://github.com/golang/go/wiki/InterfaceSlice
// var dataSlice []int = []int{1, 2, 3}
// var interfaceSlice []interface{} = dataSlice

func main() {
	s := []int{1, 2, 3}
	res1 := add(s...)

	generic := make([]interface{}, 0)
	for _, f := range s {
		generic = append(generic, f)
	}
	// 这里传参和上面注释遇到的问题一样
	res2 := add3(generic...) // slice... 展开slice

	println(res1, res2)
}

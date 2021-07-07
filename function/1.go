package main

// 匿名函数
func main() {
	fn := func() {
		println("hello world!")
	}
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 1 },
	}
	println(fns[0](100))

	// function field
	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "struct fn"
		},
	}
	println(d.fn())
}

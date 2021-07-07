package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}

// defer 后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
// 但是并没有说 struct 这里的 this 指针如何处理，通过这个例子可以看出 go 语言并没有把这个明确写出来的 this 指针当作参数来看待。
func main() {
	ts := []Test{{name: "a"}, {name: "b"}, {name: "c"}}
	for _, t := range ts {
		t2 := t
		// defer t.Close()
		defer t2.Close()
	}
}

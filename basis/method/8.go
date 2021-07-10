package main

import "fmt"

type S struct {
	T
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Printf("receiver is : %v\n", t)
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法")
}
func (t *T) testP() {
	fmt.Printf("receiver is : %v\n", *t)
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
}

// Golang 方法集：每个类型都有与之关联的方法集，这会影响到接口实现规则。
// - 类型 T 方法集包含全部 receiver T 方法
// - 类型 *T 方法集包含全部 receiver T + *T 方法
// - 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法
// - 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法
// - 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法
//
// 总结：就是如果嵌套的匿名字段类型，其中的方法会被扩展出来，和属性啥的一样，但是 receiver 还是定义时的上下文（这点和 js 很不一样）

func main() {
	s1 := S{T{1}}

	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()

	s2 := &s1
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()
}

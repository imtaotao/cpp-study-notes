package main

import "fmt"

// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。
// 换句话说，接口就是一个需要实现的方法列表。

type Sayer interface {
	say()
}

type dog struct{}

type cat struct{}

// dog 实现了 Sayer 接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat 也实现了 Sayer 接口
// 接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {

}

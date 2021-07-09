package main

import "fmt"

type Sayer interface {
	say()
}

type cat struct{}

type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

// 接口类型变量能够存储所有实现了该接口的实例
// Sayer 类型的变量能够存储 dog 和 cat 类型的变量
func main() {
	var x Sayer // 声明一个 Sayer 类型的变量 x
	a := cat{}  // 实例化一个 cat
	b := dog{}  // 实例化一个 dog

	x = a   // 可以把 cat 实例直接赋值给 x
	x.say() // 喵喵喵

	x = b   // 可以把 dog 实例直接赋值给 x
	x.say() // 汪汪汪
}

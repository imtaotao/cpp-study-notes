package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

// 接口嵌套：接口与接口间可以通过嵌套创造出新的接口
func main() {
	var x animal = cat{name: "花花"}

	x.move()
	x.say()
}

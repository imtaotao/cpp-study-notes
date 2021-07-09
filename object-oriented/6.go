package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type dog struct {
	name string
}

// 实现 Sayer 接口
func (this dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", this.name)
}

func (this dog) move() {
	fmt.Printf("%s会动\n", this.name)
}

// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
// 例如，狗可以叫，也可以动。
func main() {
	var d = dog{"旺财"}
	var s Sayer = d
	var m Mover = d

	s.say()
	m.move()
}

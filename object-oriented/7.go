package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

// dog 类型实现 Mover 接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car 类型实现 Mover 接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

// Go 语言中不同的类型还可以实现同一接口
func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}

	x = a
	x.move()

	x = b
	x.move()
}

package main

import "fmt"

type Mover interface {
	move()
}

type dog struct{}

// 值类型
// func (d dog) move() {
// 	fmt.Println("狗会动")
// }

func (d *dog) move() {
	fmt.Println("狗会动")
}

// 值接收者和指针接收者实现接口的区别
func main() {
	var x Mover
	var wangcai = dog{} // 旺财是 dog 类型
	x = wangcai         // x 不可以接收 dog 类型（此时实现 Mover 接口的是 *dog 类型，所以不能给 x 传入 dog 类型的 wangcai，此时 x 只能存储 *dog 类型的值）
	var fugui = &dog{}  // 富贵是 *dog类型
	x = fugui           // x 可以接收 *dog 类型 （如果 move 的实现是值类型，这里也是可以复制的，因为 Go 语言中有对指针类型变量求值的语法糖，dog 指针 fugui 内部会自动求值 *fugui ）
}

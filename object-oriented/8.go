package main

import "fmt"

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现 WashingMachine 接口的 dry() 方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现 WashingMachine 接口的 wash() 方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

// 一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
func main() {
	var w WashingMachine = haier{}

	w.dry()
	w.wash()
}

package main

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
}

// 对比 receiver T 和 *T 的差别
// Data: 0xc0000ae008

// Value: 0xc0000ae018
// Pointer: 0xc0000ae008

// Value: 0xc0000ae020
// Pointer: 0xc0000ae008

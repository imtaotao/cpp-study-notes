package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

// 这里接收到的 User，不管外面是传递的之中还是值，这里都是值的拷贝
func (self User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func main() {
	u := User{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)
	u.TestValue()
	u.TestPointer()
	println()

	mv := User.TestValue
	mv(u)

	mp := (*User).TestPointer
	mp(&u)

	// *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2 := (*User).TestValue
	mp2(&u)
	println()

	transform()

}

// User: 0xc000118000, {1 Tom}
// TestValue: 0xc000118030, {1 Tom}
// TestPointer: 0xc000118000, &{1 Tom}
// TestValue: 0xc000118078, {1 Tom}

// 将方法 "还原" 成函数，就容易理解了。
type Data struct{}

func (Data) TestValue() {}

func (*Data) TestPointer() {}

func transform() {
	var p *Data = nil
	p.TestPointer()

	(*Data)(nil).TestPointer() // method value
	(*Data).TestPointer(nil)   // method expression

	// p.TestValue()            // invalid memory address or nil pointer dereference

	// (Data)(nil).TestValue()  // cannot convert nil to type Data
	// Data.TestValue(nil)      // cannot use nil as type Data in function argument
}

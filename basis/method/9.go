package main

import "fmt"

// Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
// 1. `instance.method(args...)`，称为 method value
// 2. `<type>.func(instance, args...)` 称为 method expression
// 两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。
// 需要注意，method value 会复制 receiver。

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func main() {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver

	// 0xc00000c030, &{1 Tom}
	// 0xc00000c030, &{1 Tom}
	// 0xc00000c030, &{1 Tom}

	println()
	copyReceiverTest()
}

func (self User) copyTest() {
	fmt.Println(self)
}

func copyReceiverTest() {
	u := User{1, "Tom"}
	mValue := u.copyTest // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u.id, u.name = 2, "Jack"
	u.copyTest()

	mValue()
}

package main

import "fmt"

type User struct {
	Name  string
	Email string
}

// 方法
// u 是值拷贝，所以修改不会影响原来的值
func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
	u.Name += "chen"
	u.Email += "tao"
}

func main() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	fmt.Printf("%v : %v \n", u1.Name, u1.Email)

	// 指针类型调用方法
	// 即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
	fmt.Printf("%v : %v \n", u2.Name, u2.Email)
}

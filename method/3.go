package main

import "fmt"

type User struct {
	Name  string
	Email string
}

// 改为指针，就会影响原来的值
func (u *User) Notify() {
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
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
	fmt.Printf("%v : %v \n", u2.Name, u2.Email)
}

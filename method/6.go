package main

import "fmt"

type User struct {
	id   int
	name string
}

// Golang 匿名字段 ：可以像字段成员那样访问匿名字段方法，编译器负责查找
type Manager struct {
	User // 简写，里面的字段类似 js 里面的扩展符，不同的是，如果和外层又冲突，还是先拿外层的
}

func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", self, self)
}

func main() {
	m := Manager{User{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}

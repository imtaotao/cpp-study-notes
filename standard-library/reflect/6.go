package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 修改结构体值
func SetValue(o interface{}) {
	// 获取指针指向的元素
	v := reflect.ValueOf(o).Elem()

	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("tao")
	}
}

func main() {
	u := User{1, "chen", 20}

	SetValue(&u)

	fmt.Println(u)
}

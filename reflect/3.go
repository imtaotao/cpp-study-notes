package main

import (
	"fmt"
	"reflect"
)

// 反射修改值
// http://shouce.jb51.net/gopl-zh/ch12/ch12-05.html
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()

	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem() 获取地址指向的值
		// v.Elem().SetFloat(7.9)
		v.Elem().Set(reflect.ValueOf(7.9))
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

func main() {
	var x float64 = 3.4
	// 反射认为下面是指针类型，不是 float 类型
	reflect_set_value(x)
	fmt.Println("main:", x)
}

package main

import (
	"fmt"
	"reflect"
)

//反射获取 interface 值信息
func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)

	k := v.Kind()
	fmt.Println(k)

	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}

func main() {
	var x float64 = 3.4
	reflect_value(x)
}

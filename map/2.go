package main

import "fmt"

// 判断某个 key 是否在 map 里面
func main() {
	scoreMap := make(map[string]int)
	scoreMap["chen"] = 99
	scoreMap["tao"] = 100

	// Go 语言中有个判断 map 中键是否存在的特殊写法，格式如下:
	// 如果 key 存在 ok 为 true,v为对应的值；不存在 ok 为 false, v 为值类型的零值
	val, ok := scoreMap["chen"]

	chen := scoreMap["chen"]

	if ok {
		fmt.Println(val, chen)
	} else {
		fmt.Println("不存在")
	}
}

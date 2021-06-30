package main

import "fmt"

// 判断某个 key 是否在 map 里面
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["chen"] = 99
	scoreMap["tao"] = 100

	val, ok := scoreMap["chen"]

	chen := scoreMap["chen"]

	if ok {
		fmt.Println(val, chen)
	} else {
		fmt.Println("不存在")
	}
}

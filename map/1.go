package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["chen"] = 99
	scoreMap["tao"] = 100

	fmt.Println(scoreMap["chen"])
	fmt.Println(scoreMap)
	fmt.Printf("type of a:%T\n", scoreMap)

	initMap()
}

// 在声明的时候填充元素（但是值类型为 interface 的时候会报错，原因暂时不明）
func initMap() {
	userInfo := map[string]string{
		"name": "chentao",
		"age":  "23",
	}

	fmt.Println("------")
	fmt.Println(userInfo)
}

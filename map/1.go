package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["chen"] = 99
	scoreMap["tao"] = 100

	fmt.Println(scoreMap["chen"])
	fmt.Println(scoreMap)
	fmt.Printf("type of a:%T\n", scoreMap)

	map1()
}

func map1() {
	userInfo := map[string]string{
		"name": "chentao",
		"age":  "23",
	}

	fmt.Println("------")
	fmt.Println(userInfo)
}

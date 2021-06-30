package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["chen"] = 99
	scoreMap["tao"] = 100

	for k, val := range scoreMap {
		fmt.Println(k, val)
	}

	fmt.Println("--------")
	sortMap()
}

// 按照顺序遍历
func sortMap() {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	scoreMap := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("name%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)           //生成0~99的随机整数
		scoreMap[key] = value
	}

	keys := make([]string, 0, 200)

	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 按照 string 排序
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

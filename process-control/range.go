package main

// Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

// for key, value := range oldMap {
// 	newMap[key] = value
// }

// 可忽略不想要的返回值，或 "_" 这个特殊变量。

func main() {
	var s = "abc"

	for i := 0; i < 3; i++ {
		println(s[i])
	}

	println("----")

	for i := range s {
		println(string(s[i]))
	}

	println("----")

	for _, c := range s {
		println(string(c))
	}
}

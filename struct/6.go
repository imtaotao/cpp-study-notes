package main

import (
	"encoding/json"
	"fmt"
)

// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
// Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
// " `key1:"value1" key2:"value2"` "

// Student 学生
type Student struct {
	ID     int    `json:"id" db:"dbId"` //通过指定 tag 实现 json 序列化该字段时的 key
	Gender string //json 序列化是默认使用字段名作为 key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
}

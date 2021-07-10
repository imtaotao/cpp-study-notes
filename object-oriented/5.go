package main

import "fmt"

// 面试题：请问下面的代码是否能通过编译？为什么
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	// var peo People = Student{}
	var peo People = &Student{} // Speak 只接受指正类型
	fmt.Println(peo.Speak("bitch"))
}

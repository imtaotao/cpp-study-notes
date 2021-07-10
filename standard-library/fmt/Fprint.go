package main

import (
	"fmt"
	"os"
)

// Fprint 系列函数会将内容输出到一个 io.Writer 接口类型的变量 w 中，我们通常用这个函数往文件中写入内容。（只要满足io.Writer接口的类型都支持写入）
//
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

// 向标准输出写入内容
func main() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./standard-library/fmt/x.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}

	name := "chentao"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

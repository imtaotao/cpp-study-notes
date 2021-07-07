package main

import "fmt"

// 这是 js 的闭包 demo
// function a(){
// 	var i = 0;
// 	function b(){
// 		console.log(++i);
// 		document.write("<h1>"+i+"</h1>");
// 	}
// 	return b;
// }

// (function(){
// 	var c = a();
// 	c();
// 	c();
// 	c();
// 	//a(); //不会有信息输出
// 	document.write("<h1>=============</h1>");
// 	var c2 = a();
// 	c2();
// 	c2();
// })();

// 上面的 js 闭包案例用 go 重写
func a() func() int {
	i := 0
	b := func() int {
		i++
		return i
	}
	return b
}

// 闭包，递归
func main() {
	c := a()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	a()
}

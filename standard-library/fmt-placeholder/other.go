package main

import "fmt"

// '+'	总是输出数值的正负号；对 %q（%+q）会生成全部是 ASCII 字符的输出（通过转义）；
// ' '	对数值，正数前加空格而负数前加负号；对字符串采用 %x 或 %X 时（% x或% X）会给各打印的字节之间加空格
// '-'	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
// '#'	八进制数前加 0（%#o），十六进制数前加 0x（%#x）或0X（%#X），指针去掉前面的 0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的 go 字面值；
// '0'	使用 0 而不是空格填充，对于数值类型会把填充的 0 放在正负号后面；

func main() {
	s := "chentao"

	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}

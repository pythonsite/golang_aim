package main

import "fmt"

/*
变量声明的简短模式，形如：x := 100。但这种声明方式有限制：

必须使用显示初始化；
不能提供数据类型，编译器会自动推导；
只能在函数内部使用简短模式；
 */

var (
	size := 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}

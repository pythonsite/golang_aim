package main

import "fmt"

/*
不能编译通过

这里复习的知识点是：

type MyInt1 int 这行代码表示我们基于int创建了一个新类型MyInt1

type MyInt2 = int 这样代码表示我们创建了int的类型别名MyInt2

所以当我们	var i1 MyInt1 = i 去赋值的时候其实是不能通过的，因为Go是强类型语言，两者是不同的类型
而MyInt2只是int的别名，本质上还是int，所以是可以赋值
 */

type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1,i2)
}

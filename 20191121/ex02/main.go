package main

import "fmt"

/*
iota是golang语言的常量计数器,只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。


如果想要跳过一些指可以用 _

更多详细的用法可以看：https://www.cnblogs.com/zsy/p/5370052.html


 */

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
	q   // 这里会接着最开始的iota继续累加，所以这里的q的值为 6
)

func main() {
	fmt.Println(x,y,z,k,p,q)
}
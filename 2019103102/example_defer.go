package main

import "fmt"

func a() {
	i := 0
	// 因为i 是值变量，传给fmt.Println的时候传递的是一个拷贝，所以最后打印的是0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	// defer 函数的执行是后进先出的原则，所以打印的是 3,2,1,0
	for i:=0;i<4;i++ {
		defer fmt.Println(i)
	}
}

func c() (i int){
	defer func() {i++}()
	return 111
}

func main() {
	//a()
	//b()
	ret := c()
	fmt.Println(ret)
}

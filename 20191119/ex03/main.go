package main

import "fmt"

/*
new([]int) 之后的list是一个*[]int类型的指针，不能对指针执行append操作，可以使用make()初始化然后执行append
同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new()
*/

func main() {
	list := new([]int)
	list = append(list,1)
	fmt.Println(list)
}

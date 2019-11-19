package main

import "fmt"

func main() {
	s := make([]int,5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 1,2,3)
	fmt.Println(s)


	s2 := make([]int, 0)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	s2 = append(s2,1,2,3,4)
	fmt.Println(s2)
}

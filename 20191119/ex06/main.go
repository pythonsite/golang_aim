package main

import "fmt"

func main() {
	sn1 := struct {
		age int
		name string
	}{age:11, name:"aaa"}
	sn2 := struct {
		age int
		name string
	}{age:11, name:"aaa"}

	if sn1 == sn2 {
		fmt.Println("s1===s2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age:23,m: map[string]string{"a":"1"}}

	sm2 := struct {
		age int
		m map[string]string
	}{age:23,m: map[string]string{"a":"1"}}
	if sm1 == sm2 {
		fmt.Println("sm1===sm2")
	}
}
package main

/*
在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。
 */

func funcMui(x, y int) (sum int, err error) {
	return x+y,nil
}

func main() {

}

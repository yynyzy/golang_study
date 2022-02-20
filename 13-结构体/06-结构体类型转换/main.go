package main

import "fmt"

type A struct {
	Num int
}
type B struct {
	Num int
}

type C A

func main() {
	var a A
	var b B
	a = A(b) //结构体转换有要求，就是结构体的字段要完全一样（包括：名字、个数和类型）
	fmt.Println(a, b)

	//结构体之间进行 type 重定义（相当于起别名），golang让我是新的数据类型，但可以进行强转
	var c C
	c = C(a)
	fmt.Println(a, c)
}

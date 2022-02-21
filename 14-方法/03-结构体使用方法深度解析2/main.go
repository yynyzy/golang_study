package main

import "fmt"

/*
Golang中的方法作用在指定的数据类型上的(即:和指定的数据类型绑定)，因此自定义类型，
都可以有方法，而不仅仅是struct，比如int , float32等都可以有方法
*/
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写一个方法，可以改变i的值
func (i *integer) print2() {
	*i = *i + 1
}
func main() {
	var i integer = 10
	i.print2()
	i.print()
}

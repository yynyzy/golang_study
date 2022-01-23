package main

import "fmt"

func addSum(num1 int, num2 int) int {
	return num1 + num2
}

//函数是一种数据类型，可以赋值给一个变量
//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func main() {
	a := myFun
	c := a(addSum, 5, 6)
	fmt.Printf("变量a的类型是%T，c的值是%d", a, c)
}

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

type myFunType func(int, int) int

func myFun1(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	a := myFun
	c := a(addSum, 5, 6)
	a1 := myFun1
	c1 := a1(addSum, 5, 6)
	fmt.Printf("变量a的类型是%T，c的值是%d", a, c)
	fmt.Printf("变量a1的类型是%T，c1的值是%d", a1, c1)
}

/*
 Go支持自定义数据类型
基本语法： type 自定义数据类型名 数据类型 //相当取了一个别名
如：type myint int
如：type myFunType  func(int, int)int //这时 myFunType 就等价于一个函数类型 func(int, int)int
*/

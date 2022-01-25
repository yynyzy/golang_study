package main

import "fmt"

func addSum(num1 int, num2 int) int {
	return num1 + num2
}

//1.函数是一种数据类型，可以赋值给一个变量
//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

/*
 2.Go支持自定义数据类型
基本语法： type 自定义数据类型名 数据类型 //相当取了一个别名
如：type myint int
如：type myFunType  func(int, int)int //这时 myFunType 就等价于一个函数类型 func(int, int)int
*/
type myFunType func(int, int) int

func myFun1(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

/*
  3.支持对函数返回值命名
*/
func getSumandSub(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

/*
	4.GO支持可变参数
	//支持0到多个参数
	func sum(args...int) sum int{
	}

	//支持1到多个参数
	func sum(n1 int,args...int) sum int{
	}

	(1) agrs 是 slice切片，通过args[index] 可以访问到各个值
*/
func sum(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	a := myFun
	c := a(addSum, 5, 6)
	a1 := myFun1
	c1 := a1(addSum, 5, 6)
	fmt.Printf("变量a的类型是%T，c的值是%d\n", a, c)
	fmt.Printf("变量a1的类型是%T，c1的值是%d\n", a1, c1)
	x, y := getSumandSub(4, 3)
	fmt.Println("变量x，y是]n\n", x, y)

	res := sum(1, 2, 3)
	fmt.Println("4.sum求和", res)
}

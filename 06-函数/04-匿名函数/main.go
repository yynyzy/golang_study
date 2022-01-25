package main

import "fmt"

/*
 如果某个函数只希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用
*/

var (
	//
	//Fun1就是一个全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//将匿名函数 func(n1 int, n2 int) int 赋给 res1 这个变量，此时可以通过res1这个变量调用
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 3)
	fmt.Println("res1=", res1)
}

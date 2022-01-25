package main

import "fmt"

/*
	闭包就是一个函数与相关的引用环境组合的一个整体
*/

//案例：累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	//
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
}

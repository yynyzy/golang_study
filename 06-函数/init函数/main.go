package main

import "fmt"

/*
	1.每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被GO运行框架调用，init函数会在main函数前调用
	2.如果一个函数包含全局变量定义，则执行流程是：全局变量定义 ->init函数调用 ->main函数调用
	3.面试题：如果 main.go引用了 utils包，且 main.go 和 utils.go 都有变量定义、init函数时,执行流程：
	utils.go：
		变量定义【执行1】
		init函数【执行2】
	main.go:
		变量定义【执行3】
		init函数【执行4】
		main函数【执行5】

	总结；先执行引用的包的变量定义与init函数
*/
var age = test()

func test() int {
	fmt.Println("test()...")
	return 90
}

func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...", age)
}

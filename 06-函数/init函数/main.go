package main

import "fmt"

/*
	1.每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被GO运行框架调用，init函数会在main函数前调用
	2.如果一个函数包含全局变量定义，则执行流程是：全局变量定义 ->init函数调用 ->main函数调用
*/

func init() {
	fmt.Println("init函数执行")
}

func main() {
	fmt.Println("main函数执行")
}

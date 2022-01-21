package main

import "fmt"

func main() {
	//1.该区域的数据值在同一类型的范围内不断变化
	var i int = 10
	i = 50
	fmt.Println("i", i)
	//i = 1.2  //int ，不能改变数据类型

	//2.变量在同一作用域（在一个函数或代码块内）不能重名
	// var i float64 =10
	// i:=99

}

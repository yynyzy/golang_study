package main

import "fmt"

var Age int = 20 //ok
// Name:="tom" //报错
/*
	错误的原因：
		Name:="tom" 等价
		var Name string
		Name = "tom" 但 Name = "tom" 赋值语句不能在函数体外，所以错误
*/
func main() {
	fmt.Println("name", Age)
	// fmt.Println("name", Name)

}

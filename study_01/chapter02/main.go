package main

import "fmt"

// func main() {
// 	var i int 
// 	i = 10
// 	fmt.Println("i=",i)
// }

func main() {
	//golang的变量使用方式1
	//第一种:指定变量类型，声明后若不赋值,使用默认值// int的默认值是0,，其它数据类型的默认值后面马上介绍
	var i int
	fmt.Println("i=",i)
	//第二种:根据值自行判定变量类型(类型推导)
	var num = 10.11
	fmt.Println( "num=", num)
	//第三种:省略var,注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误
	//下面的方式等价 var name string = "tom"
	name := "tom"
	fmt.Println( "name=", name)
}

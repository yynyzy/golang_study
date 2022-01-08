package main

import "fmt"

// func main() {
// 	var i int 
// 	i = 10
// 	fmt.Println("i=",i)
// }



// golang的变量使用方式 1
/*
func main() {
	//第一种:指定变量类型，声明后若不赋值,使用默认值
	// int的默认值是0，其它数据类型的默认值后面马上介绍
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
*/

//该案例演示golang如何一次性声明多个变量
func main() {
	// var n1, n2,n3 int
	// fmt.Println("n1=",n1,"n2=", n2,"n3=", n3)

	//一次性声明多个变量的方式2
	var n1,name , n3 = 100,"tom", 888
	fmt.Println("n1=",n1,"name=", name,"n3=",n3)

	//一次性声明多个变量的方式3，同样可以使用类型推导
	// n1, name , n3 := 100,"tom~", 888
	// fmt.Println( "n1=",n1,"name=", name,"n3=", n3)
}
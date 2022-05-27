package main

import "fmt"

/*
	1)结构体是自定义的数据类型，代表一类事物
	2）结构体变量（实例）是具体的，实际的，代表一个具体变量
*/
//定义一个公开的结构体
type Cat struct {
	Name string
	Age  int8
}

func main() {
	var cat1 Cat
	cat1.Name = "yzy"
	cat1.Age = 18

	fmt.Printf("cat1=%v\n", cat1)
	fmt.Printf("cat1=%+v\n", cat1)
	fmt.Printf("cat1=%#+v\n", cat1)
	fmt.Printf("cat1的地址=%p", &cat1) //结构体是值类型

}

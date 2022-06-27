package main

import "fmt"

//如果结构体的字段类型是:指针，slice，和map的零值都是nil ，即还没有分配空间
//如果需要使用这样的字段,需要先make，才能使用.

//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个。

type Person struct {
	Name   string
	Age    int
	scores [2]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //切片
}

func main() {
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用 slice 和 map 一定要先赋值
	p1.slice = []int{1, 2, 3}
	fmt.Println(p1.slice)
	p1.ptr = new(int)
	fmt.Println(p1.ptr)
}

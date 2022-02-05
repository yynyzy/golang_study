package main

import "fmt"

/*
切片的使用
#方式1
第一种方式:定义一个切片，然后让切片去引用一个已经创建好的数组

#方式2
第二种方式:通过make来创建切片.
基本语法:var切片名[type = make([], len, [cap])
参数说明: type:就是数据类型  len:大小   cap:指定切片容量，可选
**/
func main() {
	var slice []int = make([]int, 5, 10)
	fmt.Println("slice的元素", slice)
	slice[1] = 10
	slice[3] = 20
	fmt.Println("slice的元素", slice)
	fmt.Println("slice的长度size=", len(slice))
	fmt.Println("slice的容量cap=", cap(slice))
}

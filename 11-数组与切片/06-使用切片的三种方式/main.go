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

方式3
第3种方式:定义一个切片，直接就指定具体数组，使用原理类似make的方式。
**/
func main() {
	/*
		1)通过make方式创建切片可以指定切片的大小和容量
		2)如果没有给切片的各个元素赋值，那么就会使用默认值[int,float=>0 string =>""  bool=>false]
		3)通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素.
	**/
	var slice []int = make([]int, 5, 10)
	fmt.Println("slice的元素", slice)
	slice[1] = 10
	slice[3] = 20
	fmt.Println("slice的元素", slice)
	fmt.Println("\nslice的长度size=", len(slice))
	fmt.Println("\nslice的容量cap=", cap(slice))

	var slice2 []string = []string{"tom", "yzy"}
	fmt.Println("slice2=", slice2)
}

package main

import "fmt"

/*

**/
func main() {
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99} //声明/定义一个切片

	//slice := intArr[1:3]
	//1. slice就是切片名
	//2. intArr[1:3]表示slice引用到intArr这个数组
	//3．引用intArr数组的起始下标为1,最后的下标为3(但是不包含3)
	slice := intArr[1:3]
	fmt.Println("intArr 的值=", intArr)
	fmt.Println("slice 的值=", slice)
	fmt.Println("slice 的长度=", len(slice))
	fmt.Println("slice 的容量=", cap(slice)) //切片的容量是动态变化的
}

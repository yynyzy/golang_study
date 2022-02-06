package main

import "fmt"

/*

**/
func main() {
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99} //声明一个数组

	//定义一个切片
	slice := intArr[1:3]
	//1. slice就是切片名
	//2. intArr[1:3]表示slice引用到intArr这个数组
	//3．引用intArr数组的起始下标为1,最后的下标为3(但是不包含3)
	fmt.Println("intArr 的值=", intArr)
	fmt.Println("slice 的元素=", slice)
	fmt.Println("slice 元素的个数=", len(slice))
	fmt.Println("slice 的容量=", cap(slice)) //切片的容量是动态变化的

	fmt.Printf("intArr[1]的地址%p\n", &intArr[1]) //切片的容量是动态变化的
	fmt.Printf("slice[0]的地址%p", &slice[0])
}

package main

import "fmt"

//指针，指针变量存的是一个地址，这个地址指向的空间存的才是值
//指针细节说明：值类型，都有对应的指针类型，稀释为 *数据类型，比如 int 对应的指针为 *int， float类对应的指针为*float（基本数据类型 int系列，float系列，bool，string，数组和结构体struct）
func main() {
	var i int = 128
	// i 在内存中的地址是,&i
	fmt.Println("i 的地址是=", &i)

	//1.ptr 是一个变量指针
	//2.ptr 的类型是 *int
	//2.ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址是=%v\n", &ptr)
	fmt.Printf("ptr 指向的值是=%v\n", *ptr)

	// 内存图
	// ptr ——> 0x0011(地址为:0x00333) ——> 128(地址为: 0x0011)
	ns := 1234.12
	var ptrs *float64 = &ns
	fmt.Printf("ns 指向的值是=%v\n", *ptrs)
}

package main

import (
	"fmt"
	"unsafe"
)

//整数类型的使用

func main() {
	var i int = 135
	fmt.Println(i)

	//测试int8的范围    -2^7~2^7-1
	// int16			-2^15~2^15-1
	// int32			-2^31~2^31-1
	// int64			-2^63~2^63-1
	var y int8 = 127
	fmt.Println(y)

	//测试一下 无符号整数 uint8 （0 ~ 2^8-1）,uint16，uint32，uint64 的范围
	var k uint8 = 255
	fmt.Println(k)

	//int（系统32位 4 字节，64位 8字节）, uint（系统32位 4 字节，64位 8字节）, rune（与 int32等价） , byte(与 Uint8等价) 的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

	//整型的使用细节
	var n1 = 100 // ? n1是什么类型
	//这里我们给介绍一下如何查看某个变量的数据类型
	//fmt.Printf()可以用于做格式化输出。
	fmt.Printf("n1的类型%T \n", n1)

	//如何在程序查看某个变量的占用字节大小和数据类型（使用较多)
	var n2 int64 = 10
	//unsafe.sizeof(n1）是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2的类型%T, n2占用的字节数是%d", n2, unsafe.Sizeof(n2))

}

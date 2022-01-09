package main

import "fmt"

//整数类型的使用

func main() {
	var i int = 135
	fmt.Println(i)

	//测试int8的范围    -2^7~2^7-1
	// int16			-2^15~2^15-1
	//int 32			-2^31~2^31-1
	//int 64			-2^63~2^63-1
	var y int8 = 127
	fmt.Println(y)

	//测试一下 无符号整数 uint8 （0 ~ 2^8-1）,unit16，uint32，uint64 的范围
	var k uint8 = 255
	fmt.Println(k)

	//int（系统32位 4 字节，64位 8字节）, uint（系统32位 4 字节，64位 8字节）, rune（与 int32等价） , byte(与 unit8等价) 的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

}

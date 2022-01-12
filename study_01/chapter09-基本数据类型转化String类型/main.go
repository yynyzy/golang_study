package main

import (
	"fmt"
	_ "unsafe"
)

func main() {
	var num1 int = 99
	var num2 float64 = 0.123
	var b bool = false
	var mychar byte = 'h'
	var str string //空的字符串

	//使用第一种方式转化为string fmt.Sprintf()

	//%q 是带双引号的%v
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str 的数据类型是%T ,值是%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str 的数据类型是%T ,值是%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str 的数据类型是%T ,值是%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str 的数据类型是%T ,值是%q\n", str, str)
}

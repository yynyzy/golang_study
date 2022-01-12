package main

import (
	"fmt"
	"strconv"
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

	//第二种方式 strcon 函数
	var num3 int = 66
	var num4 float64 = 0.667
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str 的数据类型是%T ,值是%q\n", str, str)

	// 'f':格式    10：小数点保留10位   64；小数是64位
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str 的数据类型是%T ,值是%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str 的数据类型是%T ,值是%q\n", str, str)
}

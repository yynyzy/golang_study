package main

import (
	"fmt"
	"strconv"
)

//演示golang中string转成基本数据类型
func main() {
	var str string = "true"
	var b bool

	// b._ = strconv.ParseBool(str)
	//说明:
	// 1. strconv.ParseBool(str）函数会返回两个值（value bool, err error)
	// 2．因为我只想获取到value bool ,不想获取err所以我使用_忽略

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "9111111111"
	var b2 int64
	//ParseInt 返回的是 int64 类型，所以b2要为 int64 才能接收（如果要变为 int8等，通过 int8(b2)等手动转换）
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 是检查你想要转换的数是否超过了 bitSize 的限制（64位范围）
	b2, _ = strconv.ParseInt(str2, 10, 64) // 10：转为10进制
	fmt.Printf("b2 type %T b2=%v\n", b2, b2)

	var str3 string = "123.456"
	var f1 float64
	//也是函数默认返回的是 float64位
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}

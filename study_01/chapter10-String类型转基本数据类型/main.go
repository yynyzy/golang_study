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
	//ParseInt 返回的是 int64 类型，所以b2要为int64 才能接受
	b2, _ = strconv.ParseInt(str2, 10, 64) // 10：转为10进制
	fmt.Printf("b2 type %T b=%v\n", b2, b2)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1.统计字符串的长度，按字节 len(str)
	str := "hello北"
	//golang的编码统一为utf-8（ascil的字符（字母和数字）占一个字符，汉字占用3个字节）
	fmt.Println("str的长度=", len(str))

	//2.字符串遍历，同时处理有中文的问题 []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//3.字符串转整数：
	// n, err := strconv.Atoi("hello") //invalid syntax
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("错误转换", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

}

package main

import (
	"fmt"
)

func main() {
	/*
		switch分支结构
		switch 表达式{
			case 表达式1，表达式2，...：语句块1

			case 表达式1，表达式2，...：语句块1
			//这里可以有多个case

			default：
			语句块
		}
	*/

	var key byte
	fmt.Println("请输入字符a或b")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	default:
		fmt.Println("星期天")
	}
}

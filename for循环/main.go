package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("你好，尚硅谷", i)
	}
	// for循环的第二种写法
	j := 1        //循环变量初始化
	for j <= 10 { //循环条件
		fmt.Println("你好，尚硅谷~", j)
		j++ //循环变量迭代
	}

	//for循环的第三种写法,这种写法通常会配合break使用
	k := 1
	for { //这里等价 for ;;{
		if k <= 10 {
			fmt.Println("ok~", k)
		} else {
			break //跳出循环
		}
		k++
	}

	//字符串遍历方式一
	var str string = "握手,yzy"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	//字符串遍历方式二——for range
	//Unicode遍历
	str = "abc~12握手"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c \n", index, val)
	}

}

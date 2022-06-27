package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("你好", i)
	}
	// for循环的第二种写法
	j := 1       //循环变量初始化
	for j <= 5 { //循环条件
		fmt.Println("你好", j)
		j++ //循环变量迭代
	}

	//for循环的第三种写法,这种写法通常会配合break使用
	k := 1
	for { //这里等价 for ;;{
		if k <= 3 {
			fmt.Println("ok~", k)
		} else {
			break //跳出循环
		}
		k++
	}

	//字符串遍历方式一传统方式
	/*
		有中文会乱码，因为传统的对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码对应的是 3个字节。所以输出一个字节就会乱码。
		解决方法：使用 []rune切片
	*/
	// var str string = "握手,yzy"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i])
	// }
	var str string = "握手,yzy"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	//字符串遍历方式二——for range
	//Unicode遍历
	str = "abc~12握手"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c \n", index, val)
	}

}

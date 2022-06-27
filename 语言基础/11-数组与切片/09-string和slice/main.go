package main

import "fmt"

/*
1) string 底层是一个byte数组，因此 string 也可以进行切片处理
2) string 是不可变的，也就是说不能通过 str[0]='z' 方式来修改字符串
3) 如果需要修改字符串，可以先将string -> []byte 或者 []rune ->修改 ->重写转成string
**/
func main() {
	var str string = "hello@yzy"
	//使用切片获取到yzy
	slice := str[6:]
	fmt.Println("slice=", slice)

	//如果需要修改字符串，可以先将string -> []byte 或者 []rune ->修改 ->重写转成string
	// arr1 := []byte(str)
	// arr1[0] = 'z'
	// str = string(arr1)
	// fmt.Println("str=", str)

	//z转成[]byte 后，可以处理英文和数字，但不能处理中文
	//原因是 []byte 字节来处理，而一个汉字，是三个字节，因此会出现乱码
	//解决办法是将 string 转成 []rune 即可，因为 []rune是按字符处理，兼容汉字
	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1)
	fmt.Println("str=", str)
}

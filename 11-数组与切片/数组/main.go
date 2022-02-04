package main

import "fmt"

/**
数组的定义
var 数组名 【数组大小】数据类型
var a [5]int
赋初值 a[0]=1 a[1]=30
*/

func main() {
	var intArr [3]int32 //int占8个字节(int64)
	//当我们定义完数组后,其实数组的各个元素有默认值
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0]地址%p\nintArr[1]地址%p\nintArr[2]地址%p",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
}

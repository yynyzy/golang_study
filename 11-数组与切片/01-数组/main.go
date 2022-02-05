package main

import "fmt"

/**
数组的定义
var 数组名 【数组大小】数据类型
var a [5]int
赋初值 a[0]=1 a[1]=30
*/

func main() {
	var intArr [3]int //int占8个字节(int64)
	//当我们定义完数组后,其实数组的各个元素有默认值 0
	fmt.Println(intArr)
	//1.数组的地址可以通过数组名来获取&intArr
	//2数组的第一个元素的地址，就是数组的首地址
	fmt.Printf("intArr的地址=%p intArr[0]地址%p\nintArr[1]地址%p\nintArr[2]地址%p",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
}

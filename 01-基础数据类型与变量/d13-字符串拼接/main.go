package main

import "fmt"

func main() {
	const a string = "YYN"
	const b string = "Yzy"

	// 1.使用"+"会对字符串进行遍历，计算并开辟一个新的空间来存储原来的两个字符串
	const c1 = a + b
	fmt.Println("c1=", c1)

	//2.fmt.Sprintf 由于采用了接口参数，必须要用反射获取值，因此有性能损耗。
	c2 := fmt.Sprintf("%s%s", a, b)
	fmt.Println("c2=", c2)
}

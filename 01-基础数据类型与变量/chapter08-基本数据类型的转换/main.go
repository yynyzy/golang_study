package main

import "fmt"

func main() {
	var i int32 = 100
	var i1 float32 = float32(i)
	var i2 int64 = int64(i)
	fmt.Printf("i1的值是%v，数据类型是%T \n", i1, i1)
	fmt.Printf("i2的值是%v，数据类型是%T \n", i2, i2)

	// var n1 int32 =127
	// var n2 int64
	// var n3 int8
	// n2 = n1 + 20    		//int32 ---> int64错误   (不能将int32赋值给int64，除非经过转化再赋值)
	// n2 = int64(n1) + 20
	// n3 = int8(n1) + 20
	// fmt.Println("n2=", n2,"n3=", n3)

	var n1 int32 = 12
	var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 //【编译通过，n3就是int8类型，127在运算过程中没有超出 int8范围，但结果溢出不是127+12,】
	// n3 = int8(n1) + 128 //【编订不过，n3就是int8类型，128在加法过程中就超过了 int8的范围，报错】
	fmt.Println(n3, n4)

}

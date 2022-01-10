package main

import (
	"fmt"
)

//演示golang中小数类型使用
func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)
	var num1 float32 = -0.00889
	var num2 float64 = -7809656.09
	fmt.Println("num1=", num1, "num2=", num2)

	//尾数部分可能丢失,造成精度损失。-123.0000901
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	//Golang的浮点型默认声明为float64类型
	var num5 = 1.1
	fmt.Printf("num5的数据类型是%T \n", num5)

	//十进制数形式:如:5.12      .512(必须有小数点)
	num6 := 5.12
	num7 := .123 //=>0.123
	fmt.Println("num6=", num6, "num7=", num7)

	//科学计数法形式
	num8 := 5.1234e2  // ? 5.1234 * 10的2次方
	num9 := 5.1234e-2 // ? 5.1234 * 10的2次方
	fmt.Println("num8=", num8, "num9=", num9)

}

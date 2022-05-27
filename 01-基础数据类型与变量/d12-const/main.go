package main

import "fmt"

/*
·常量使用const修改
·常量在定义的时候，必须初始化
·常量不能修改
·常量只能修饰bool、数值类型(int, float系列)、 string 类型
·语法: const identifier [type] = value
·举例说明，看看下面的写法是否正确:
	const name = "tom' 			//ok
	const tax float64 = 0.8  	//ok
	const a int					//ok
	const b= 9/3				//ok   (如果9变成一个变量会报错)
	constc = getVal()			//err
*/
func main() {
	//常量在定义的时候，必须初始化给一个值
	const num int = 10
	fmt.Println(num)

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		a2     = iota
		b2, c2 = iota, iota //按行来的，在一行不会递增
	)
	fmt.Println(a2, b2, c2)
}

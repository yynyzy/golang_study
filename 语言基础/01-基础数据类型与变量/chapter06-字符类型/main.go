package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	//当我们输出byte值，就是输出来对应的码值
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)

	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c", c1, c2)

	// var c3 byte = '北' //overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c c3的对应码值=%d\n", c3, c3)

	//可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode 字符
	var c4 int = 22269 // 22269 ->‘国’120->'x'
	fmt.Printf("c4=%c\n", c4)
	//字符类型是可以进行运算的，相当于一个整数,运输时是按照码值运行
	var n1 = 10 + 'a' //10 + 97 = 1o7
	fmt.Println("n1=", n1)

}

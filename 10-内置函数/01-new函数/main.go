package main

import "fmt"

//内建函数new分配内存。其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值的指针。

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T , num1的值=%v , num1的地址%v\n", num1, num1, &num1)

	num2 := new(int) //*int
	//num2的类型T=>*int
	//num2的值= 0xc00000a0d0(这个地址是系统分配)
	//num2的地址= 0xc000006030(这个地址是系统分配)
	// num2指向的值默认 = 0
	fmt.Printf("num2的类型%T , num2的值=%v , num2的地址%v,num2这个指针，指向的值=%v\n", num2, num2, &num2, *num2)
	*num2 = 100 // num2指向的值改为 100
	fmt.Printf("num2指向的值=%v", *num2)

}

package main

import "fmt"

// 10.2.2方法的声明和调用

type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test()", p.Name)
}

func (p Person) add(n1, n2 int) int {
	res := n1 + n2
	return res
}

/*
	方法的调用和传参机制原理:(重要!)
	说明:方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参也传递给方法(如果变量是值类型，则进行值拷贝，如果是引用类型，则进行地址拷贝)。下
*/

//1) func(p Person) test{}   表示 Person 结构体有一方法，方法名为test
//2) (p Person)体现 test方法是和 Person 类型绑定的
//3） test 方法只能通过 Person 类型的变量来调用，而不能直接调用，也不能使用其他类型来调用
func main() {
	var p Person

	p.Name = "Tom"
	// test()  错误的调用,必须使用 p.test() 的用法

	p.test()                      //调用方法，把 p的副本传进去
	fmt.Println("main()", p.Name) //"Tom"（和 p.test() 不影响）

	res := p.add(1, 2)
	fmt.Println("res=", res)
}

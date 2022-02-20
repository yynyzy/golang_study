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

//对上面的语法的说明

//1) func(p Person) test{}   表示 Person 结构体有一方法，方法名为test
//2) (p Person)体现 test方法是和 Person 类型绑定的
//3） test 方法只能通过 Perso 类型的变量来调用，而不能直接调用，也不能使用其他类型来调用
func main() {
	var p Person

	p.Name = "Tom"
	p.test()                      //调用方法，把 p 传进去
	fmt.Println("main()", p.Name) //值传递所以不影响

	// test()   //错误的调用
}

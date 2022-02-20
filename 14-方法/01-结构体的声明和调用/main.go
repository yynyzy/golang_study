package main

import "fmt"

// 10.2.2方法的声明和调用
type Person struct {
	Num int
}

func (p Person) test() {
	fmt.Println("test()", p.Num)
}

//对上面的语法的说明

//1) func(p Person) test{}   表示 Person 结构体有一方法，方法名为test
//2) (p Person)体现 test方法是和 Person 类型绑定的
func main() {
	var p Person
	p.test()
}

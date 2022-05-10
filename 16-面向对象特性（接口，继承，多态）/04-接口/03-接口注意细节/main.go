/*
	·注意事项和细节
	1)接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)
	2)接口中所有的方法都没有方法体,即都是没有实现的方法。
	3)在Golang中，-个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现
	了该接口。
	4)一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
*/
package main

import "fmt"

type AInterface interface {
	say()
}
type Stu struct {
}

func (stu Stu) say() {
	fmt.Println("stu say()")
}

type integer int

func (i integer) say() {
	fmt.Println("integer Say i =", i)
}
func main() {
	var stu Stu //结构体变量,实现了Say() 实现了AInterface
	var a AInterface = stu
	a.say()
	var i integer = 10
	var b AInterface = i
	b.say() // integer Say i = 10
}

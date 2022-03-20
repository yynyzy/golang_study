package main

import (
	"fmt"
	"reflect"
)

/*
基本介绍
	1)反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)，类别(kind)
	2)如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法)
	3)通过反射，可以修改变量的值，可以调用关联的方法。
	4)使用反射，需要import ("reflect")
*/
func test(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(b)
	fmt.Println("rType=", rVal)
	n1 := 10 + rVal.Int()
	fmt.Println("n1=", n1)

}
func main() {
	var num int = 8
	test(num)
}

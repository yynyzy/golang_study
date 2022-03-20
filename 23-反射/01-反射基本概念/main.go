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
	//获取reflect.TypeOf
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//获取reflect.ValueOf
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v,rVal的类型是=%T\n", rVal, rVal) //rVal的类型是 reflect.Value

	//将rVal 转换之后才能与 10 相加，不然类型不一样
	n1 := 10 + rVal.Int()
	fmt.Println("n1=", n1)

	//将 rVal 转成 interface{},最后重新转成 int 类型
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}
func main() {
	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作

	var num int = 8
	test(num)
}

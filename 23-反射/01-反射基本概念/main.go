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
func test01(b interface{}) {
	//通过反射获取传入变量的 type
	rType := reflect.TypeOf(b)
	fmt.Println("type=", rType)

	//获取reflect.ValueOf
	rVal := reflect.ValueOf(b)
	fmt.Printf("val=%v,val的类型是=%T\n", rVal, rVal) //rVal的类型是 reflect.Value

	//获取变量的类型 kind
	// rType.kind()
	// rVal.kind()
	fmt.Printf("kind1=%v,kind2=%v\n", rType.Kind(), rVal.Kind())

	//将rVal 转换之后才能与 10 相加，不然类型不一样
	n1 := 10 + rVal.Int()
	fmt.Println("n1=", n1)

	//将 rVal 转成 interface{},最后重新转成 int 类型
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}

//以下是对 结构体的反射 处理
func test02(b interface{}) {
	//获取reflect.TypeOf
	rType := reflect.TypeOf(b)
	fmt.Println("type=", rType)

	//获取reflect.ValueOf
	rVal := reflect.ValueOf(b)
	fmt.Printf("val=%v,val的类型是=%T\n", rVal, rVal)

	//获取变量的类型 kind
	// rType.kind()
	// rVal.kind()
	fmt.Printf("kind1=%v,kind2=%v\n", rType.Kind(), rVal.Kind())

	//将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	//也可以使用 switch 的断言形式来做的更加灵活
	v, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", v.Name)
	}

}

type Monster struct {
	Name string
	Age  int8
}
type Student struct {
	Name string
	Age  int8
}

func main() {
	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作

	var num int = 8
	test01(num) //对基本数据类型

	var stu = Student{"yzy", 18}
	// var mon = Monster{"yzy", 18}
	test02(stu) //对 结构体
	// test02(mon)
}

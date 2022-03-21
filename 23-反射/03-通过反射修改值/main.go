package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	rVal := reflect.ValueOf(b)
	//传入地址时，它的kind 也是一个地址

	fmt.Printf("rVal的类型是=%v\n", rVal.Kind()) //rVal的类型是=ptr
	rVal.Elem().SetInt(10)                   //地址要用 .Elem() 来修改
}
func main() {
	var num int = 10
	test(&num)
	fmt.Println("num=", num)
}

package main

import "fmt"

func main() {
	var i int32 = 100
	var i1 float32 = float32(i)
	var i2 int64 = int64(i)
	fmt.Printf("i1的值是%v，数据类型是%T \n", i1, i1)
	fmt.Printf("i2的值是%v，数据类型是%T \n", i2, i2)

}

package main

import "fmt"

//定义一个公开的结构体
type Cat struct {
	Name string
	Age  int8
}

func main() {
	var cat1 Cat
	fmt.Println(cat1)

}

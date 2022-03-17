package main

import "fmt"

type Cat struct {
	Name string
	Age  int8
}

func main() {
	cat := Cat{"yyn", 18}
	//定义一个可以存放任意数据类型的管道
	var allchain = make(chan interface{}, 3)
	allchain <- 10
	allchain <- "yzy"
	allchain <- cat
	fmt.Printf("allchain的类型是=%T,值是=%v\n", allchain, allchain)

	<-allchain
	<-allchain
	cat1 := <-allchain
	//直接取出来的是 interface类型的 cat1，直接cat.Name 编译不通过
	// fmt.Printf("cat1的Name=%v", cat1.Name)
	// fmt.Printf("cat1=%v,cat1的类型是=%T", cat1, cat1)

	a := cat1.(Cat)
	fmt.Printf("a的Name=%v\n", a.Name)
	fmt.Printf("a=%v,cat1的类型是=%T\n", a, a)
}

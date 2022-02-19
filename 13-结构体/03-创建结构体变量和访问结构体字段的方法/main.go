package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式一

	// 方式二
	p2 := Person{"mary", 20}
	// p2.Name = "tom"
	// p2.Age = 18
	fmt.Println(p2)

	//方式3-&
	//案例:var person *Person = new (Person)
	var p3 *Person = new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	// (*P3).Name = "smith" 也可以这样写  p3.Name = "smith"
	//原因:go的设计者为了程序员使用方便，底层会对 p3.Name = "smith"进行处理
	//会给p3加上取值运算(*P3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "john"

	(*p3).Age = 30
	p3.Age = 100
	fmt.Println(*p3)

}

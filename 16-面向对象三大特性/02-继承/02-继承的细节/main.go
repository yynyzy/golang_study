package main

import("fmt")


type A struct {
Name string
age int
}

type B struct {
Name string
Score float64
}

type C struct {
	A
	B
	//Name string
}
type D struct {
	a A
}

func main() {
	var c C
	//如果c没有Name字段，而 A 和 B 有 Name, 这时就必须通过指定匿名结构体名字来区分
	//所以 c.Name 就会包编译错误,这个规则对方法也是一样的
	c.A.Name = "tom" 
	fmt.Println("c")


	//如果D中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
	//比如d.a . Name
	var d D
	d.a.Name = "jack"
}
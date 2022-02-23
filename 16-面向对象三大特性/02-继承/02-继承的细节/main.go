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
type Goods struct{
	Name string
	Price float64
}
type Brand struct{
	Name string
	Address string
}

type TV struct{
	Goods
	Brand
}

type TV2 struct{
	*Goods
	*Brand
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


	test2()
	
}

func test2 (){
	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}, }
tv2 := TV{
Goods{
Price : 5000.99,
Name : "电视机002" ,
},
Brand{
Name : "夏普",
Address :"北京" ,
},

}
fmt. Println("tv", tv)
fmt . Println("tv2", tv2)


tv3 := TV2{ &Goods{"电视机003", 7000.99}, &Brand{"创维", "湖南"}, }
fmt.Println("tv3",*tv3.Goods,*tv3.Brand)

}
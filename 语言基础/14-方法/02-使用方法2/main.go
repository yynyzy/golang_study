package main

import "fmt"

type Circle struct {
	radius float64
}

//为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	//因为c是指针，因此我们标准的访问其字段的方式是（*c).radius
	//return 3.14 * (*c).radius * (*c).radius
	// (*c).radius 等价 c.radius
	c.radius = 10
	fmt.Printf(" c 是 *Circle指向的地址%p\n", c)
	return 3.14 * c.radius * c.radius
}
func main() {
	//创建一个circle变量
	//c.radius = 4.0l/ res := c.area()
	// fmt.Print1n("面积是=", res)
	//创建一个circle变量
	var c Circle
	c.radius = 7.0

	res2 := c.area2()
	//等价于 res2 :=(&c).area2()
	//编译器底层做了优化 (&c).area2()等价 c.area()
	//因为编译器会自动的给加上 &c
	fmt.Println("面积=", res2)
	fmt.Println("c.radius = ", c.radius) //10 因为传的是地址
	fmt.Printf("mian c 结构体变量的地址是%p\n", &c)
}

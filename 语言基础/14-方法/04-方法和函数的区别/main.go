package main

import "fmt"

/*
1)调用方式不一样
		函数的调用方式:            函数名(实参列表)
		方法的调用方式:            变量.方法名(实参列表)
2)对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
3)对于方法（如struct的方法)，接收者为值类型时，可以直接用指针类型的变量调用方法(但不是真正的引用，不影响外面),反过来同样也可以。

总结:
	1)不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定.
	2)如果是和值类型，比如(p Person)，则是值拷贝，如果和指针类型，比如是(p *Person)则是地址拷贝。

*/
type Person struct {
	Name string
}

func test01(p Person) {
	fmt.Println("teste1() =", p.Name) //tom
}
func test02(p *Person) {
	fmt.Println("teste2() =", p.Name) //tom
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("teste3() =", p.Name) //jack
}
func (p Person) test04() {
	p.Name = "marry"
	fmt.Println("teste3() =", p.Name) //marry
}
func main() {
	p := Person{"tom"}
	test01(p)
	test02(&p)
	p.test03()
	fmt.Println("main() test03() p.name=", p.Name) // tom

	(&p).test03()                                  //从形式上是传入地址，但是本质仍然是值拷贝
	fmt.Println("main() test03() p.name=", p.Name) // tom

	(&p).test04()                                  //这是标准的引用传递的方式
	fmt.Println("main() test04() p.name=", p.Name) // marry
}

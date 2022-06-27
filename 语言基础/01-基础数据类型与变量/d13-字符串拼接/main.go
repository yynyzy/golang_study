package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	const a string = "YYN"
	const b string = "Yzy"

	// 1.使用"+"会对字符串进行遍历，计算并开辟一个新的空间来存储原来的两个字符串
	const c1 = a + b
	fmt.Println("c1=", c1)

	//2.fmt.Sprintf 由于采用了接口参数，必须要用反射获取值，因此有性能损耗。
	c2 := fmt.Sprintf("%s%s", a, b)
	fmt.Println("c2=", c2)

	//3.strings.Builder：
	//用WriteString()进行拼接，内部实现是指针+切片，同时String()返回拼接后的字符串，它是直接把[]byte转换为string，从而避免变量拷贝。
	var c3 strings.Builder
	c3.WriteString(a)
	c3.WriteString(b)
	C3 := c3.String()
	fmt.Println("C3=", C3)

	//4.bytes.Buffer是一个一个缓冲byte类型的缓冲器，这个缓冲器里存放着都是byte。
	var c4 bytes.Buffer
	c4.WriteString(a)
	c4.WriteString(b)
	fmt.Println("c4=", c4.String())

	//5.strings.join也是基于strings.builder来实现的,并且可以自定义分隔符
	c5 := strings.Join([]string{a, b}, " ")
	fmt.Println("c5=", c5)
}

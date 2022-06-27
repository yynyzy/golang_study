package main

import (
	"fmt"
)

//演示golang中string类型使用
func main() {
	//string的基本使用
	var address string = "北京长城110 hello wor1d!"
	fmt.Println(address)

	//字符串一旦赋值了，字符串就不能修改了:在Go中字符串是不可变的
	//var str ="hello"
	//str[0] = 'a'
	//这里就不能去修改str的内容，即go中的字符串是不可变的。

	//字符串的两种表示形式:
	//(1)双引号,会识别转义字符
	//(2)反引号，以字符串的原生形式输出，包括换行和特殊字符,可以实现防止攻击
	str2 := "abc\nabc"
	fmt.Println(str2)

	str3 := `
	package main
	import "fmt"
	func main() {
		var n2 = 200
		fmt.Println("n2=",n2)
	}
	`
	fmt.Println(str3)
}

package main

import (
	"errors"
	"fmt"
)

// 自定义错误
// Go程序中，也支持自定义错误，使用 errors.New 和 panic 内置函数。
// 1)errors.New("错误说明")，会返回一个error类型的值，表示一个错误
// 2) panic内置函数,接收一 个interface{}类型的值 (也就是任何值了)作为参数。可以接收 error 类型的变量，输出错误信息，并退出程序.

//函数去读取以配置文件init. conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	//读取
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误~")
	}
}

func test() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
}
func main() {
	test()
	fmt.Println("test()函数下面的代码 执行")
}

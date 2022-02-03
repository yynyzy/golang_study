package main

import (
	"fmt"
	"time"
)

// 1)Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally这种处理。
// 2) Go中引入的处理方式为:defer, panic, recover
// 3)这几个异常的使用场景可以这么简单描述。Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内 置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res的值为", res)
}

func main() {
	test()
	for {
		fmt.Println("下面的代码执行")
		time.Sleep(time.Second)
	}

}

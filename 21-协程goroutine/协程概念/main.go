package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Go协程和Go主线程
1) Go主线程(有程序员直接称为线程/也可以理解成进程):一一个Go线程上，可以起
多个协程，你可以这样理解，协程是轻量级的线程。

2) Go协程的特点
	·有独立的栈空间
	·共享程序堆空间
	·调度由用户控制
	·协程是轻量级的线程
*/

/*
	1.在go主线程中开启一个 goroutine，每隔一秒输出一个 helloworld
	2.在go主线程中也每隔一秒输出一个 helloworld
*/

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println("test() hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go test() //开启协程
	for i := 0; i < 5; i++ {
		fmt.Println("main() hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

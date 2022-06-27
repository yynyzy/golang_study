package main

import (
	"fmt"
	"runtime"
)

func say(s string) {

	for i := 0; i < 2; i++ {

		fmt.Println(s)

	}

}

func main() {

	go say("java")

	for i := 0; i < 2; i++ {
		runtime.Gosched() //把当前的CPU 时间线让出去
		fmt.Println("golang")
	}
	fmt.Println("hello")

	/*
		java
		java
		golang
		golang
		hello
	*/
}

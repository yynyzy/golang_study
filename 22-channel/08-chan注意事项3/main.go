package main

import (
	"fmt"
	"time"
)

/*
	协程中使用 recover ，解决协程中出现 panic，导致程序崩溃
*/
func sleep() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("sleep()", i)
	}
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("test()协程发生错误err=%d", err)
		}
	}()
	var intmap map[int]string
	intmap[0] = "hello"
}
func main() {
	go sleep()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main()", i)
		time.Sleep(time.Second)
	}
}

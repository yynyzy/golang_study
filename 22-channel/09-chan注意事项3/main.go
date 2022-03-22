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
	//使用 recover 可以不会让一个出错的协程影响整个程序
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("test()协程发生错误err=%d", err)
		}
	}()

	var intmap map[int]string //因为没有 make 空间，所以报错
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

package main

import (
	"fmt"
	"sync/atomic"
)

var i int32 = 100

//atomic 提供的原子操作能够保证任意时刻只有一个 goroutine 对变量进行操作，善用 atomic 能够避免程序中出现大量的锁操作
func add() {
	atomic.AddInt32(&i, 1)
}
func sub() {
	atomic.AddInt32(&i, -1)
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	fmt.Println("i=", i)
}

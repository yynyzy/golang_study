package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i:%v", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i:%v", i)

	var j int64 = 200
	atomic.AddInt64(&j, 1)
	fmt.Printf("j:%v", j)
}
func test_read_write() {
	var i int32 = 100
	atomic.LoadInt32(&i) //原子读，保证读的时候不会有其他线程对其进行修改
	fmt.Printf("i:%v\n", i)

	atomic.StoreInt32(&i, 222) //原子写，保证写的时候不会有其他线程对其进行修改
	fmt.Printf("i:%v\n", i)
}
func main() {
	var i int32 = 100
	res := atomic.CompareAndSwapInt32(&i, 100, 1) //将 &i与 100进行比较，相同就替换为 1，返回 T or F。（原子操作为了不让与 100 比较时 100 被修改）
	fmt.Printf("i:%v\n", i)
	fmt.Println("bool=", res)
}

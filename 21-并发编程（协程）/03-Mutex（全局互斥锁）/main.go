package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	计算1～100 各个数的阶乘，把结果放入到 map 中

	问题：1.会出现两个协程同时修改一个map【key】的 值
		 2.无法确定程序需要延迟多少秒来让协程执行
*/
var (
	mapTable = make(map[int]int, 10)
	//添加一个全局的互斥锁
	lock sync.Mutex
)

func test(limit int) {
	res := 1
	for i := 1; i <= limit; i++ {
		res *= i
	}
	//加锁,防止同时对mapTable操作
	lock.Lock()
	mapTable[limit] = res
	//解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//设置来防止协程还未全部开启或完成任务，主线程就执行完成，导致协程也一并销毁
	time.Sleep(time.Second)

	for k, v := range mapTable {
		fmt.Printf("map[%d]=%v\n", k, v)
	}
}

package main

import (
	"fmt"
	"time"
)

/*
	计算1～100 各个数的阶乘，把结果放入到 map 中

	问题：会出现两个协程同时修改一个map【key】的 值
*/
var (
	mapTable = make(map[int]int, 10)
)

func test(limit int) {
	res := 1
	for i := 1; i <= limit; i++ {
		res *= i
	}
	mapTable[limit] = res
}

func main() {
	for i := 1; i <= 100; i++ {
		go test(i)
	}
	//设置来防止协程还未全部开启或完成任务，主线程就执行完成，导致协程也一并销毁
	time.Sleep(time.Second * 20)

	for k, v := range mapTable {
		fmt.Printf("map[%d]=%v\n", k, v)
	}
}

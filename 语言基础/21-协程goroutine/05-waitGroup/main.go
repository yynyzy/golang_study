package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func showMsg(i int) {
	defer wg.Done()
	//等价于 defer wg.Add(-1)   任务完成 -1

	fmt.Printf("i=%v\n", i)
}
func main() {
	for i := 0; i < 10; i++ {

		//启动一个协程来执行
		go showMsg(i)
		wg.Add(1) //增加一个任务
	}

	wg.Wait()
	//主线程
	fmt.Println("end...")
}

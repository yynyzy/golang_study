package main

import "fmt"

/*
	向管道写入50个数据，同时读取
*/
func writeChan(inChan chan int) {
	for i := 1; i <= 50; i++ {
		inChan <- i
		fmt.Println("inChan 写入数据=", i)
	}
	close(inChan)
}
func readChan(inChan chan int, exitChan chan bool) {
	for {
		v, ok := <-inChan
		if !ok {
			break
		}
		fmt.Println("inChan 读取数据=", v)
	}

	exitChan <- true
	close(exitChan)
}
func main() {
	var inChan = make(chan int, 50)
	var exitChan = make(chan bool, 1)

	go writeChan(inChan)
	go readChan(inChan, exitChan)
	//通过exitchan 来不让 主线程提前完成
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

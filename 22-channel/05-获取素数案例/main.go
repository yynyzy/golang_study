package main

import "fmt"

/*
	获取1～8000中所有的素数
*/

//写入8000个数到 inChan 管道中
func writeData(inChan chan int) {
	for i := 1; i <= 8000; i++ {
		inChan <- i
	}
	close(inChan)
}
func findPrime(inChan chan int, primeChan chan int, exitChan chan bool) {

	for {
		var flag bool
		num, ok := <-inChan
		if !ok { //intchan 取不到
			break
		}
		flag = true
		//判断是否是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个 primeChan协程取不到数据，关闭")
	//这里不能关闭 primeChan

	exitChan <- true

}

func main() {
	var inChan = make(chan int, 8000)
	var primeChan = make(chan int, 8000)
	var exitChan = make(chan bool, 4)
	go writeData(inChan)

	for i := 0; i < 4; i++ {
		go findPrime(inChan, primeChan, exitChan)
	}

	go func(primeChan chan int) {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitChan 中取出 4个结果，说明四个协程全部跑完，可以关闭 primeChan
		close(primeChan)
	}(primeChan)

	//遍历 primeChan 取出素数
	for v := range primeChan {
		fmt.Println("primeChan素数=", v)
	}
	fmt.Println("主线程结束！")
}

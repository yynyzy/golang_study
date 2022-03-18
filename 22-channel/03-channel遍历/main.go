package main

import (
	"fmt"
)

func main() {
	var intchan chan int = make(chan int, 3)
	intchan <- 10
	intchan <- 20
	//关闭管道后无法在进行写操作，只能读操作
	close(intchan)
	// intchan <- 300

	//管道关闭后，可以读数据
	n1 := <-intchan
	fmt.Println("n1=", n1)

	var intchan2 = make(chan int, 20)
	for i := 1; i <= 20; i++ {
		intchan2 <- i * 2
	}

	// 遍历管道不能使用普通的for循环
	// for i:=0;i<len(intchan2);i++{

	// }

	//在进行遍历时，管道没有关闭，会报错，deadlock
	close(intchan2)
	for v := range intchan2 {
		fmt.Println("intchan2 v=", v)
	}
}

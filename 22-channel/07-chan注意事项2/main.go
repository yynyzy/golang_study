package main

import "fmt"

func main() {
	var intchan = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intchan <- i
	}

	var stringchan = make(chan string, 5)
	for i := 1; i <= 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock

	//问题，在实际开发中，可能我们不好确定什么时候关闭管道。
	//可以使用 select 方式解决
	for {
		select {
		//注意：这里如果intchan一直没有关闭，不会一直阻塞而 deadlock
		//，会自动到下一个case 匹配
		case v := <-intchan:
			fmt.Println("从intchan中取出数据", v)
		case v := <-stringchan:
			fmt.Println("从stringchan中取出数据", v)
		default:
			fmt.Println("没有数据可以取了，退出")
			return
		}
	}
}

package main

import "fmt"

/*
	1)channel本质就是一个数据结构-队列
	2）数据是先进先出
	3）线程安全，多 goroutine 访问时，不需要加锁，就说channel 本身就是线程安全的
	4）channel 是有类型的，一个 string 的 channel 只能存放 string 类型的数据
*/
/*
 1）channel 是引用类型
 2）channel 必须初始化才能使用，及先make后，才能使用
 3）管道是用类型的 intchan，只能写入 int类型数据
*/
func main() {
	var mychan chan int = make(chan int, 3)
	fmt.Printf("mychan=%v,mychan自己本身的地址是%v\n", mychan, &mychan)

	mychan <- 10
	num := 20
	mychan <- num

	//管道有长度和容量，管道的容量是固定的，不能存放超过容量的数据
	fmt.Printf("mychan的 len=%v,cap=%v\n", len(mychan), cap(mychan))
	mychan <- 40

	var num1 int
	num1 = <-mychan

	num2 := <-mychan
	num3 := <-mychan
	fmt.Printf("mychan中取出来的数%v,%v,%v\n", num1, num2, num3)
	//在没有使用协程的情况下，管道数据已经取完，在获取就会报错  deadlock!死锁
	// num4 := <-mychan
	// fmt.Println("num4=", num4)
}

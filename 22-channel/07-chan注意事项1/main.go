package main

func main() {
	//默认情况下，管道是双向的
	// var intchan chan int

	//声明为只写管道
	// var intchan chan<- int

	//声明为只读管道
	// var intchan <-chan int

	//例子

	var intchan chan int
	go send(intchan)    //send 函数只接受 只写属性的管道
	go receive(intchan) //receive 函数只接受 只读属性的管道
}
func send(ch chan<- int) {
	//只写操作
}
func receive(ch <-chan int) {
	//只读操作
}

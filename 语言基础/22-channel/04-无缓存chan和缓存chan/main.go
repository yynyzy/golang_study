package main

/*
1.无缓存channel
channel是否有缓存最根本的不同是阻塞策略的不同。
·无缓存的channel会在读写两端产生阻塞即：当对一个无缓存的channel读时，若数据尚未到达，则读协程阻塞；当对一个无缓存的channel写时，若数据尚未消费，则写协程阻塞。

·基于此，无缓存的channel也成为同步channel

·需要特别注意的是，当数据尚未被消费时，接收者接收数据发生在唤醒发送者协程之前！
*/

/*
2.有缓存channel
带缓存的channel实际上是一个阻塞队列。对队列的写总发生在队尾，队列满时写协程阻塞；对队列的读总发生在队尾，队列空时读协程阻塞。

内置函数cap(channel)可以读取队列的总长度。
内置队列len(channel)可以读取队列的有效长度（已使用长度）。

有缓存的channel的最大意义在于生产和消费可以解耦。
*/

func main() {
	//这种容量为0的channel就是无缓存channel
	// channel_test := make(chan string)
	// channel_test := make(chan string,0)

	//有缓存channel
	// channel_test := make(chan string,1)
}

package main

/*
	1)channel本质就是一个数据结构-队列
	2）数据是先进先出
	3）线程安全，多 goroutine 访问时，不需要加锁，就说channel 本身就是线程安全的
	4）channel 是有类型的，一个 string 的 channel 只能存放 string 类型的数据
*/

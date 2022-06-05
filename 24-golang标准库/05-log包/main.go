package main

import "log"

/*
1og包中有3个系列的日志打印函数。分别Drint系列、Danic 系列、fatal 系列。

函数系列					    作用
print						单纯打印日志
panic						打印日志，抛出panic异常
fatal						打印日志，强制结束程序(os.Exit()),defer 函数不会执行
*/

func test1() {
	log.Print("my log")
	log.Printf("my log%d", 100)
}

// panic
func test2() {
	defer log.Println("panic 结束后执行")
	log.Print("my log")
	log.Panic("my panic")
	log.Panicln("end...")
}

// fatal
func test3() {
	defer log.Println("Fatal 结束后执行")
	log.Print("my log")
	log.Fatal("my Fatal")
	log.Panicln("end...")
}
func main() {
	// test1()
	// test2()
	test3()
}

package main

import "fmt"

/*
	在函数中，程序员经常需要创建资源（比如：数据库连接，文件句柄，锁等），为了在函数执行完毕后，及时的释放资源，Go的设计者提供 defer（延时机制）
*/
func sum(n1 int, n2 int) int {
	//当执行到 defer 时，暂时不执行，将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok1 n1=", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	sum(10, 20) //30 -> 20 -> 10
}

package main

import "fmt"

/*
	在函数中，程序员经常需要创建资源（比如：数据库连接，文件句柄，锁等），为了在函数执行完毕后，及时的释放资源，Go的设计者提供 defer（延时机制）
*/

/*
用途：
1)在golang编程中的通常做法是，创建资源后，比如(打开了文件，获取了数据库的链接，或者是锁资源),可以执行defer file.Close()、defer connect.Close()
2)在defer后，可以继续使用创建的资源
3)当函数完毕后，系统会依次从 defer 栈中，取出语句，关闭资源.
4)这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心。
*/
func sum(n1 int, n2 int) int {
	//当执行到 defer 时，暂时不执行，将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	//在 defer 将语句放入栈中的时候。也会将相关的值同时拷贝入栈
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok1 n1=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	sum(10, 20) //30 -> 20 -> 10
}

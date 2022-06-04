package main

import (
	"fmt"
	"os"
)

/*
流:数据在数据源(文件)和程序(内存)之间经历的路径
输入流:数据从数据源(文件)到程序(内存)的路径
输出流:数据从程序(内存)到数据源(文件)的路径
*/

func main() {
	//file 可以叫 file对象，file句柄，file指针
	//open 打开文件只能只读，openfile 可写可读
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Open file error!")
	} else {
		fmt.Printf("%p\n", file) //是个指针
	}
	err = file.Close()
	if err != nil {
		fmt.Println("文件打开失败")
	}
}

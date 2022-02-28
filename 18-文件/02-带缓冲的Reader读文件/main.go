package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../test.txt")
	if err != nil {
		fmt.Println("打开文件失败！")
	}
	defer file.Close() //打开文件后要及时关闭文件句柄，defer在最后调用
	reader := bufio.NewReader(file)
	/*
		const (
			defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件末尾，读到文件末尾就推出
			break
		}
		//输出内容
		fmt.Print(str)
	}
}

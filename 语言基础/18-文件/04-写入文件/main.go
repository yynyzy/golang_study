package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filepath := "/Users/yanyinuo/go/src/golang_study/18-文件/test03.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open File Error%v\n", err)
		return
	}
	//及时关闭file 句柄
	defer file.Close()

	str := "写入的内容哈哈哈\n"
	//使用带缓存的 *writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//写入五次
		writer.WriteString(str)
	}
	/*
		因为writer是带缓存的，因此在调用writerString的方法时，其实内容写入的是内存，
		需要调用 Flush 方法将数据真正写入到文件中
	*/
	writer.Flush()

}

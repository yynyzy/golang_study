package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		ioutil 一次性读取文件到内存，前提是文件不大的情况下
		ioutil.ReadFile 不需要open和close ，因为 ReadFile 已经把它们全部完成了
	*/
	content, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		fmt.Println("打开文件失败！")
	} else {
		fmt.Println(content) //[]byte
		fmt.Printf("%v", string(content))

	}
}

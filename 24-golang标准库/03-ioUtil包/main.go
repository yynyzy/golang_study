package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
封装一些实用的 I/O 函数
名称                 |                     作用
ReadAll				|		读取数据，返回读到的字节 slice
ReadDir				|		读取一个目录，返回目录入口数组 []os.Filelnfo
ReadFile			|		读一个文件，返回文件内容 (字节slice)
WriteFile			|		根据文件路径，写入字节 slice
TempDir				|		在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径
TempFile			|		在一个目录中创建指定前缀名的临时文件，返回os.File
*/

//一次读取文件
func test_ReadAll() {
	// r := strings.NewReader("Yzy Hello")
	f, _ := os.Open("go_study\\24-golang标准库\\03-ioUtil包\\a.txt") //返回的 file 结构体实现了 Reader 接口
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err = %v\n", err)
	} else {
		fmt.Printf("string(b)= %v\n", string(b))
	}
}

//读目录
func test_ReadDir() {

	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Printf("v.Name()= %v\n", v.Name())
	}

}
func main() {
	test_ReadAll()
	test_ReadDir()
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client Dial err=%v\n", err)
		return
	}
	//功能一：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)

	//从终端读取一行用户输入，并准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error ReadString err=", err)
	}
	//再将line发送给 服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("Error writing err=", err)
	}
	fmt.Printf("客户端发送了%d字节的数据，并退出", n)
}

package main

import (
	"fmt"
	"net"
)

//处理和客户端的通信
func process(conn net.Conn) {
	//等待客户端发送信息
	defer conn.Close()
	buf := make([]byte, 8096)
	for {
		fmt.Println("等待客户端发送的数据")
		_, err := conn.Read(buf[:4]) //n是指读到了多少个字节
		if err != nil {
			fmt.Println("conn.Read() err=", err)
			return
		}
		fmt.Println("读到的buf=", buf[:4])
	}
}
func main() {
	fmt.Println("服务端监听8889端口")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("服务端等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)
	}
}

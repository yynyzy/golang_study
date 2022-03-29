package main

import (
	"fmt"
	"net"
)

//处理和客户端的通信
func process(conn net.Conn) {
	//等待客户端发送信息
	defer conn.Close()
	for {
		buf := make([]byte, 8096)
		fmt.Println("等待客户端发送的数据")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read() err=", err)
			return
		}
		fmt.Println("读到的buf")
	}
}
func main() {
	fmt.Println("服务端监听8889端口")
	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	for {
		fmt.Println("服务端等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)
	}
}

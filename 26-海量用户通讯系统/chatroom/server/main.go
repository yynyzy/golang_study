package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//等待客户端发送信息
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

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务端开始监听~")
	//net.Listen("tcp", "0.0.0.0:8888")
	//  tcp 表示使用网络协议
	// "0.0.0.0:8888" 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close() //延时关闭 listen

	//等待客户端来链接我
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept() err=%v\n", err)
		} else {
			fmt.Printf("Accept() success conn=%v\n", conn)
		}
		//这里准备一个协程，为客户端服务
	}

	// fmt.Printf("listen=%v\n", listen)
}

package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	//这里我们循环的接受客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//等待客户端通过 conn 发送信息
		//2.如果客户端没有write【发送】，那么协程在这阻塞
		// fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从 conn 读取
		if err == io.EOF {
			fmt.Printf("客户端退出err=%v\n", err)
			return //!!!
		}
		//显示客户端发送的内容到服务终端
		fmt.Println("接受到客户端发送的数据：", string(buf[:n]))

	}
}
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
			fmt.Printf("Accept() success conn=%v,客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备一个协程，为客户端服务
		go process(conn)
	}

	// fmt.Printf("listen=%v\n", listen)
}

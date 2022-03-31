package main

import (
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/server/utils"
	"io"
	"net"
)

//处理和客户端的通信
func process(conn net.Conn) {
	//等待客户端发送信息
	defer conn.Close()
	for {
		tf := &utils.Transfer{Conn: conn}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err=", err)
			if err == io.EOF {
				fmt.Println("客户端退出,服务器此 connect 也正常关闭")
				return
			} else {
				fmt.Println("readPkg err=", err)
				return
			}
		}
		p := &Processor{
			Conn: conn,
		}
		err = p.ServerProcesMes(&mes)
		if err != nil {
			return
		}
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

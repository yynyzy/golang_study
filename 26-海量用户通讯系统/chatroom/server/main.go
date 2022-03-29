package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"io"
	"net"
)

//处理和客户端的通信
func process(conn net.Conn) {
	//等待客户端发送信息
	defer conn.Close()
	for {
		mes, err := readPkg(conn)
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
		fmt.Println("mes=", mes)
	}
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	//conn.Read() 在conn没有关闭的情况下才会阻塞，如果客户端关闭要

	_, err = conn.Read(buf[:4]) //读取消息的头四个字节，获取应得到的长度
	//这里的buf[:4] 是指：从 conn套接字中读取了 4个字节的数据到 buf中
	//返回的n是指读到了多少个字节
	if err != nil {
		return
	}

	//根据 buf[:4] 转成一个 uint32类型
	var pkglen uint32 = binary.BigEndian.Uint32(buf[0:4])

	//根据  从套接字中读取 pkglen 长度的数据到 buf 中去
	n, err := conn.Read(buf[:pkglen])
	if n != int(pkglen) || err != nil {
		return
	}

	//根据 pkglen 读取相应长度的数据-> 反序列化为 message.Message
	err = json.Unmarshal(buf[:pkglen], &mes)
	if err != nil {
		return
	}

	return
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

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

//编写一个函数serverProcessLogin函数， 专门处理登录请求
func serverprocessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码...
	//1. 先从mes中取出mes.Data ，并直接反序列化成LoginMes
	var loginMes message.Login_Message
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json. Unmarshal fail err=", err)
		return
	}
	//如果用户id= 100,密码=123,认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
	} else {
	}
	return
}

//编写一个 ServerProcessMes 函数
//功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcesMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.Login_Mes_Type:
		//处理登陆
		fmt.Println("")
	case message.Register_Mes_Type:
		//处理注册
		fmt.Println("")

	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
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

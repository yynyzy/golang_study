package main

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"golang_study/26-海量用户通讯系统/chatroom/common/utils"
	"io"
	"net"
)

//处理和客户端的通信
func process(conn net.Conn) {
	//等待客户端发送信息
	defer conn.Close()
	for {
		mes, err := utils.ReadPkg(conn)
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
		err = serverProcesMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

//编写一个 ServerProcessMes 函数
//功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcesMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.Login_Mes_Type:
		//处理登陆
		err = serverProcessLogin(conn, mes)
	case message.Register_Mes_Type:
		//处理注册
		fmt.Println("")

	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//编写一个函数serverProcessLogin函数， 专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码...
	//1. 先从mes中取出mes.Data ，并直接反序列化成LoginMes
	var loginMes message.Login_Message
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json. Unmarshal fail err=", err)
		return
	}
	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.Login_Res_Mes_Type

	//再声明一个 loginResMes
	var loginResMes message.Login_Response_Message

	//如果用户id= 100,密码=123,认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法,返回500
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册后再登陆..."
	}
	//3将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6.发送
	err = utils.WritePkg(conn, data)
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

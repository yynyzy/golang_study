package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"net"
)

func login(UserId int, UserPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889") //连接到服务端
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.Login_Mes_Type
	var loginMes message.Login_Message
	loginMes.UserId = UserId
	loginMes.UserPwd = UserPwd

	//将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//把data 赋值给 mes.Data 字段
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//注意网络传输的丢包问题
	//将要发送的数据的长度先发送过去，告诉服务器接下来数据有多长，再发送数据
	//服务端在接受到数据后，会判断数据是不是等于预期的长度，如果不一样，说明丢包，需要解决
	var pkglen uint32
	pkglen = uint32(len(data))
	var byteArr [4]byte
	binary.BigEndian.PutUint32(byteArr[0:4], pkglen)
	//发送长度
	n, err := conn.Write(byteArr)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err", err)
		return
	}
	fmt.Println("客户端发送信息的长度ok")
}

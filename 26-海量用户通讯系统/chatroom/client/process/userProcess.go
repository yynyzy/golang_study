package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/client/utils"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(UserId int, UserPwd string, UserName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889") //连接到服务端
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.Register_Mes_Type     //message 类型
	var RegisterMes message.Register_Message //message 的data
	RegisterMes.User.UserId = UserId
	RegisterMes.User.UserPwd = UserPwd
	RegisterMes.User.UserName = UserName

	//将 loginMes 序列化
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		fmt.Println("json.Marshal(RegisterMes) err=", err)
		return
	}
	//把data 赋值给 mes.Data 字段
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Register WritePkg(data) err=", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Register ReadPkg err=", err)
		return
	}
	//将 mes 的Data部分反序列化成 RegisterResMes
	//服务端向客户端返回一个响应消息
	var RegisterResMes message.Register_Response_Message
	err = json.Unmarshal([]byte(mes.Data), &RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功，请重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(RegisterResMes.Error)
		os.Exit(0)
	}
	return
}

func (this *UserProcess) Login(UserId int, UserPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889") //连接到服务端
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.Login_Mes_Type  //message 类型
	var loginMes message.Login_Message //message 的data
	loginMes.UserId = UserId
	loginMes.UserPwd = UserPwd

	//将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
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
	var pkglen uint32 = uint32(len(data))
	var buf [4]byte //4字节的buf用来储存最多32位无符号位的数据
	binary.BigEndian.PutUint32(buf[:4], pkglen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err", err)
		return
	}
	fmt.Printf("客户端发送信息的长度=%d,内容是=%s\n", pkglen, string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(data) fail err", err)
		return
	}
	tf := &utils.Transfer{Conn: conn}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	//将 mes 的Data部分反序列化成LoginResMes
	//服务端向客户端返回一个响应消息
	var loginResMes message.Login_Response_Message
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		//可以显示当前在线用户列表,遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			fmt.Println("用户id:\t", v)
		}
		fmt.Print("\n\n")

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯.如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端.
		go serverProcessMes(conn)
		//1.显示我们的登录成功的菜单[循环]..
		for {
			showMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

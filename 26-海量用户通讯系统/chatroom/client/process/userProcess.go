package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/client/utils"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) login(UserId int, UserPwd string) (err error) {

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
	var pkglen uint32 = uint32(len(data))
	var buf [4]byte //4字节的buf用来储存最多32位无符号位的数据
	binary.BigEndian.PutUint32(buf[:4], pkglen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err", err)
		return
	}
	// fmt.Printf("客户端发送信息的长度=%d,内容是=%s\n", pkglen, string(data))

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
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

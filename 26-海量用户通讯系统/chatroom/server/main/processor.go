package main

import (
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	processes "golang_study/26-海量用户通讯系统/chatroom/server/process"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//编写一个 ServerProcessMes 函数
//功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) ServerProcesMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.Login_Mes_Type:
		//处理登陆
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.Register_Mes_Type:
		//处理注册
		fmt.Println("")

	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

package model

import (
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"net"
)

//设为全局变量，方便客户端使用
type CurUser struct {
	Conn net.Conn
	message.User
}

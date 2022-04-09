package process

import "golang_study/26-海量用户通讯系统/chatroom/common/message"

//客户端维护的用户状态map

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

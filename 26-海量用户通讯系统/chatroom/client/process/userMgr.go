package process

import (
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
)

//客户端维护的用户状态map

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//在客 户端显示当前在线的用户
func outputonlineUser() {
	//遍历一 把onlineUsers
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		//如果不显示自己，
		fmt.Println("用户id:\t", id)
	}
}

//处理返回的 Notify_User_Status_Mes
func updateUsersStatus(Notify_User_Status_Mes *message.Notify_User_Status_Mes) {
	user, ok := onlineUsers[Notify_User_Status_Mes.UserId]
	if !ok {
		user = &message.User{
			UserId: Notify_User_Status_Mes.UserId,
		}
	}
	user.UserStatus = Notify_User_Status_Mes.Status
	onlineUsers[Notify_User_Status_Mes.UserId] = user
	fmt.Println("updateUsersStatus Map=", onlineUsers)
}

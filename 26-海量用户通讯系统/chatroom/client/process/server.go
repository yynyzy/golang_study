package process

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/client/utils"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"net"
	"os"
)

func showMenu() {
	fmt.Println("-------恭喜xxx登录成功--------")
	fmt.Println("-------1.显示在线用户列表--------")
	fmt.Println("----------2.发送消息---------")
	fmt.Println("----------3.信息列表---------")
	fmt.Println("----------4.退出系统---------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputonlineUser()
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")

	}
}

//和服务器保持通讯，接受服务器推送的消息
func serverProcessMes(Conn net.Conn) {
	//创建一个 Transfer 实例，不停的读取服务器发送的信息
	tf := &utils.Transfer{Conn: Conn}

	for {
		fmt.Println("客户端正在读取服务器推送的信息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err=", err)
			return
		}

		switch mes.Type {
		case message.Notify_User_Status_Mes_Type:
			//1.有用户上线的消息推送
			//2.把这个用户的消息状态更新到客户端的 map[int]User 中
			var notifyUserStatusMes message.Notify_User_Status_Mes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUsersStatus(&notifyUserStatusMes)
		default:
			fmt.Println("服务端返回了未知的消息类型")
		}
		//如果读取到消息，进行进行下一步逻辑
		fmt.Printf("mes=%v", mes)
	}
}

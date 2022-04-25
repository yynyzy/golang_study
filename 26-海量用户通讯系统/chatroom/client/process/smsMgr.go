package process

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}
	//显示消息
	info := fmt.Sprintf("用户id:\t%d对大家说:\t%s\n", smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println("")
}

package process

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/client/utils"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMessage(Content string) (err error) {
	//创建一个mes
	var mes message.Message
	mes.Type = message.SmsMes_Type

	//创建一个 smsMes实例
	var smsMes message.SmsMes
	smsMes.Content = Content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMessage json.Marshal(smsMes) err=", err.Error())
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMessage json.Marshal(mes) err=", err.Error())
		return
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMessage tf.WritePkg(data) err=", err.Error())
		return
	}
	return
}
